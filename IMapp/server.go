package main

import (
	"fmt"
	"net"
	"sync"
	_"time"
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
	outMessage := "User [" + user.Addr + "]" + user.Name + ":" + msg
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
			cli.UserCh <- msg
		}
		s.mapLock.Unlock()
	}
}

// offer service to the connected client
func (s *Server) Handler(conn net.Conn) {
	fmt.Println("Client connected successfully")
	user := NewUser(conn)
	
	// add user to the online user map
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()
	// broadcast the user online status
	s.BroadCast(user, "is online now. Welcome!")
	// defer conn.Close()
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

		// respond to the client
		go s.Handler(conn)

	}
}