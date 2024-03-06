package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"bufio"
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

// menu utility
func (client *Client) menu() bool {
	var mode string

	fmt.Println("Please select mode:")
	for i := 0; i < 4; i++ {
		fmt.Println(i, modeMap[i])
	}

	fmt.Scanln(&mode)
	modeInt, err := strconv.Atoi(mode)
	if err != nil {
		fmt.Println("invalid input")
		return false
	}
	if modeInt >= 0 && modeInt <= 3 {
		client.Mode = modeInt
		fmt.Println("Now you are in mode:", modeMap[modeInt])
		return true
	} else {
		fmt.Println("invalid input")
		return false
	}
}

// deal with response from the server
func (client *Client) DealResponse() {
	// [continuously]listern the response from the server and print to console
	io.Copy(os.Stdout, client.Conn)
}

// update username
func (client *Client) UpdateUsername() bool {
	var newName string
	if client.Name != "" {
		fmt.Println("Your current username is", client.Name)
	}
	fmt.Println("Please input new username:")
	fmt.Scanln(&newName)
	sendMsg := "rename:" + newName + "\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("error:", err)
		return false
	}
	client.Name = newName
	return true
}

// public message
func (client *Client) PublicMessage() bool {
	var msg string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input message: type exit [to exit the chat]")
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error:", err)
		return false
	}	
	msg = strings.TrimRight(msg, "\n")

	for msg != "exit" {
		if msg != "" {
			sendMsg := msg + "\n"
			_, err := client.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("error:", err)
				return false
			}
		}
		msg = ""
		fmt.Println("Please input message: type exit [to exit the chat]")
		msg, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("error:", err)
			return false
		}	
		msg = strings.TrimRight(msg, "\n")	
	}
	return true
}

// query online users
func (client *Client) QueryOnlineUsers() bool {
	sendMsg := "search\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("error:", err)
		return false
	}
	return true
}

// private message
func (client *Client) PrivateMessage() bool {
	var (msg string
		toName string
	)
	client.QueryOnlineUsers() // query online users
	fmt.Println("Please input the username to send message:")
	fmt.Scanln(&toName)
	toName = strings.Trim(toName, " ")
	if toName == "" {
		fmt.Println("invalid input")
		return false
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input message: type exit [to exit the chat]")
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error:", err)
		return false
	}
	msg = strings.TrimRight(msg, "\n")	

	for msg != "exit" {
		if msg != "" {
			sendMsg := "to:" + toName + ":" + msg + "\n"
			_, err := client.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("error:", err)
				return false
			}
		}
		msg = ""
		fmt.Println("Please input message: type exit [to exit the chat]")
		msg, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("error:", err)
			return false
		}
		msg = strings.TrimRight(msg, "\n")
	}
	return true
}

// main goroutine
func (client *Client) Run() {
	for client.Mode != 0 {
		for client.menu() { // like while
			switch client.Mode {
			case 1:
				client.PublicMessage()
			case 2:
				client.PrivateMessage()
			case 3:
				client.UpdateUsername()
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
	go client.DealResponse() // start a goroutine to deal response from the server.
	
	client.Run()
}