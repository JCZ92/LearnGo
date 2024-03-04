package main
import (
	"fmt"
	"time"
	"sync"
)

func Routine2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		fmt.Println("from other Goroutine", i)
		time.Sleep(time.Millisecond * 500)
	}
}