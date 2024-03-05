package main
import "fmt"


func main() {
	// send only channel, it's a property not a type
	var intchan chan<- int 
	intchan = make(chan int, 3)
	intchan <- 10
	fmt.Println("send only channel", intchan)
	// fmt.Println(<-intchan) . // invalid

	// receive only channel, it's a property not a typel
	var intchan2 <-chan int
	intchan2 = make(chan int, 3)
	fmt.Println("receive only channel", intchan2)
	// intchan2 <- 10 // invalid

}