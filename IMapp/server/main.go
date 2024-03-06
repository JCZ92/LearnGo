package main

var serverPort int = 9999
var serverIP string = "localhost"
var bufferSizeForMessageFromClient int = 4096
var timeoutLogOffInSec int = 1000

func main() {
	server := NewServer(serverIP, serverPort)
	server.Start()
}