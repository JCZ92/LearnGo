package main

import (
	"fmt"
)

// channel is a queue under the hood, is goroutine safe, channel has its own containing datatype
// channel is a typed conduit through which you can send and receive values with the channel operator <-
// channel <- value sends a value down the channel

func main() {
	//declare a channel
	// var intChan chan int
	int1 := 10
	ptr := &int1
	c := make(chan int, 5) // initialize a channel of capa = 3, return a memory addresss, but not a pointer!!! though you are seeing an address %T
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", ptr)
	fmt.Printf("%v\n", ptr)

	// store data to channel
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	
	// channel has properties: len and cap
	fmt.Println(len(c), cap(c))

	// read data from channel one by one, popping from the underlying queue
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(len(c), cap(c))

	// close the channel and you can no longer write to the channel, but you can still read from the channel
	close(c)
	fmt.Println(<-c)
	// c <- 1 will panic

	// iterate over the channel, have to close it before iteration, otherwise will [deadlock]
	for v := range c {
		fmt.Print(v, "\t")
	}
	
}
