package main

import (
	// "fmt"
	// "time"
	"fmt"
	"sync"
)

type Locker interface{
	Lock()
	Unlock()
}


func main() {
	
	// go Routine() //start a goroutine
	// note that once the main goroutine exits, the other goroutine will also exit

	// for i := 0; i < 2; i++ {
	// 	fmt.Println("from main Goroutine", i)
	// 	time.Sleep(time.Millisecond * 500)
	// }


	// to fix the above issue, we can use a waitgroup

	// var wg sync.WaitGroup
	// for i := 0; i < 3; i++ {	
	// 	wg.Add(1)
	// 	go Routine(&wg)
	// }
	// wg.Wait()


	// multiple goroutines can access the same variable, prevent race condition by adding mutex lock
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// num := 0
	// wg.Add(2)
	// go Add(&num, &wg, &mu)
	// go Sub(&num, &wg, &mu)
	// fmt.Println(num)
	// wg.Wait()


	// use RWMutex for situation where read are much more than write
	var wg sync.WaitGroup
	
	// var mu sync.Mutex
	var mu sync.RWMutex // you can toggle the type of mutex easily
	num := 0
	wg.Add(2)
	go Add(&num, &wg, &mu) // has to use &mu here, since Mutex defines Lock/Unlock using a pointer receiver, not a value receiver
	go Sub(&num, &wg, &mu)
	fmt.Println(num)
	wg.Wait()

}
