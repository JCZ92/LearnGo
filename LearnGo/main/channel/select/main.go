package main
import (
	"fmt"
	"time"
)

func main() {
	intchan := make(chan int, 2)
	go func() {
		time.Sleep(time.Second)
		intchan <- 10
	}()

	stringchan := make(chan string, 2)
	go func() {
		time.Sleep(time.Second * 2)
		stringchan <- "hello"
	}()
	

	// select, execute the first case that is not empty, otherwise, execute the second case
	select {
	case num := <- intchan:
		fmt.Println("num = ", num)
	case str := <- stringchan:
		fmt.Println("str = ", str)
	case <- time.After(time.Second * 3): // this is good
		fmt.Println("timeout")
	// default:
	// 	fmt.Println("default")
	}

}