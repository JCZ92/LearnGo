package main

import (
	"fmt"
	"net"
	"sync"
	"time"
	"io"
)

type Server struct {
	Ip string
	Port int

	// add a map for online users
	OnlineMap map[string]*User
	mapLock sync.RWMutex // protect the map

	// add a channel for message
	Message chan string
}

// define a factory
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}
}

// broadcast message[user is now online] to all online users
// push the message to server's channel
func (s *Server) BroadCast(user *User, msg string) {
	outMessage := "[" + user.Addr + "]" + user.Name + msg
	s.Message <- outMessage
}

// continuously listen the message channel and push the message to all online users 
func (s *Server) PushMessageToOnlineUsers() {
	for {
		msg, ok := <-s.Message // will block if no message is received
		if !ok {
			fmt.Println("Message channel closed")
			break
		}
		// send the message to all online users
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.SendMessage(msg)
		}
		s.mapLock.Unlock()
	}
}

// offer service to the connected client
func (s *Server) Handler(conn net.Conn) {
	fmt.Println("Client connected successfully")
	user := NewUser(conn, s)
	// refactor online step as user method 
	user.Online()
	// create a channel to determine if user is active
	isActive := make(chan bool)

	// broadcast the message from the user to all online users
	go func() {
		buf := make([]byte, 4096) // creat a buffer to receive messages from client
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			}
			if (err == io.EOF) {
				return
			}
			// get the message from the client
			msg := string(buf[:n-1])
			if msg != "" { // only broadcast the message if it is not empty
				go user.HandleMsg(msg)
			}
			isActive <- true // user is active
		}
	}()

	// Timeout enforced logoff
	for {
		select { 
			// execute the first case that is not empty
			// in this case, every select will wait at most 5 seconds, if in this 5 seconds, an isActive is received, will reset the timer
			// if timeout, the user is forced to log off
		case <-isActive:
			// do nothing, user is active
		case <-time.After(time.Second * 300):
			user.SendMessage("You are being logged off due to inactivity in one second")
			time.Sleep(time.Second)
			//close the connection
			conn.Close()
			return
		}
	}
}

// start a server
func (s *Server) Start() {
	fmt.Println("Server starts, listening", fmt.Sprintf("%s:%d", s.Ip, s.Port))

	// listen a socket
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		fmt.Println("Server stops")
		return
	}

	//close the socket
	defer listener.Close()

	// start a goroutine to listen and push message to all online users
	go s.PushMessageToOnlineUsers()

	// continuously accept clients and then respond
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		// handle to the client requests
		go s.Handler(conn)

	}
}