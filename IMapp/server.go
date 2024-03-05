package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip string
	Port int
}

// define a factory
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

// offer service to the connected client
func (s *Server) Handler(conn net.Conn) {
	fmt.Println("Client connected successfully")
	defer conn.Close()
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