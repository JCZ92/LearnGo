package main
import (
	"fmt"
	"time"
	"sync"
)

func Routine(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("from other Goroutine", i)
		time.Sleep(time.Millisecond * 500)
	}
}