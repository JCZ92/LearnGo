package main
import (
	"sync"
)


func Sub(num *int, wg *sync.WaitGroup, lock Locker) {
	defer wg.Done()
	_, okay := lock.(*sync.RWMutex);
	for i := 0; i < 50; i++ {
		if  okay {
			lock.(*sync.RWMutex).RLock() // you can also call method onto a type assertion
		} else {
			lock.Lock()
		}
		*num--;
		if  okay {
			lock.(*sync.RWMutex).RUnlock()
		} else {
			lock.Unlock()
		}
	}
}