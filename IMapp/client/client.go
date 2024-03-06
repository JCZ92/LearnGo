package main

import(
	"fmt"
	"net"
)

type Client struct {
	ServerIp string
	ServerPort int
	Name string
	Conn net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	// possible to give name?
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.Conn = conn

	return client;
}

func main() {
	client := NewClient("127.0.0.1", 9999)
	if client == nil {
		fmt.Println("fail to create a client")
		return
	}
	fmt.Println("client is created")
	select {}
}