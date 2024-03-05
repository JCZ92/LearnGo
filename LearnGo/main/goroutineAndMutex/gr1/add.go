package main
import (
	"sync"
)


func Add(num *int, wg *sync.WaitGroup, lock Locker) {
	defer wg.Done()
	L, okay := lock.(*sync.RWMutex);
	for i := 0; i < 50; i++ {
		if  okay {
			L.RLock()
		} else {
			lock.Lock()
		}
		*num++;
		if  okay {
			L.RUnlock()
		} else {
			lock.Unlock()
		}
	}
}