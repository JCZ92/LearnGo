package main

import (
	"fmt"
	"sync"
	"time"
)

// producer: write data to a channer
// consumer: read data from a channel
// once all data in the channel is read, the channel is closed and exit the main goroutine
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		intChan <- i
		fmt.Println("writeData", i)
		time.Sleep(time.Millisecond * 300)
	}
	close(intChan)
}

func readData(intChan chan int) {
	defer wg.Done()
	for {
		v, ok := <-intChan 
		// ok is a boolean value that indicates whether the channel is closed or not. If the channel is closed, ok will be set to false. If the channel is not closed, ok will be set to true.
		if !ok {
			break
		}
		fmt.Println("readData", v)
		time.Sleep(time.Millisecond * 100)
	}

}

// wait group to wait for all goroutines to finish before exiting the main goroutine
var wg sync.WaitGroup
func main() {

	ch := make(chan int,  10)
	wg.Add(2)
	go writeData(ch)
	go readData(ch)
	wg.Wait()
}
