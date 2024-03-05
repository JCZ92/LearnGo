package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("client start...")
	// protocol and port
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	fmt.Println("client start success!", conn)

	for {
		// get input from stdin
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err, exit!")
			return
		}

		//send data to server and then exit
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println("write string err, exit!")
			return
		}
	}

	// defer conn.Close()
}