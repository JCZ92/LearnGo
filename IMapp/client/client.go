package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp string
	ServerPort int
	Name string
	Conn net.Conn
	Mode int
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
	client.Mode = 9999
	return client;
}
// global var
var (
	serverIp string
	serverPort int
	modeMap map[int]string
)

func (client *Client) menu() bool {
	var mode int

	fmt.Println("Please select mode:")
	for i := 0; i < 4; i++ {
		fmt.Println(i, modeMap[i])
	}

	fmt.Scanln(&mode)
	if mode >= 0 && mode <= 3 {
		client.Mode = mode
		fmt.Println("Now you are in mode:", modeMap[mode])
		return true
	} else {
		fmt.Println("invalid input")
		return false
	}
}

func (client *Client) Run() {
	for client.Mode != 0 {
		for client.menu() { // like while
			switch client.Mode {
			case 1:
				fmt.Println("broadcast message to all")
			case 2:
				fmt.Println("private message")
			case 3:
				fmt.Println("change username")
			case 0:
				return
			}
		}
	}
}

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server ip, default is 127.0.0.1")
	flag.IntVar(&serverPort, "port", 9999, "set server port, default is 9999")
	modeMap = make(map[int]string)
	modeMap[1] = "broadcast message to all"
	modeMap[2] = "private message"
	modeMap[3] = "change username"
	modeMap[0] = "exit"
}

func main() {
	flag.Parse() // parse the parameters from command line.
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("fail to create a client")
		return
	}
	fmt.Println("client is created")
	client.Run()
}