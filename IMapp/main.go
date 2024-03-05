package main

func main() {
	server := NewServer("localhost", 9999)
	server.Start()
}