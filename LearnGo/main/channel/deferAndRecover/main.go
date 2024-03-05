package main

import (
	"fmt"
	"time"
)

func main() {
	
	go func() {
		defer func() { // define inside the goroutine
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("panic")
	}()

	go func() {
		fmt.Println("hello")
	}()

	time.Sleep(time.Second) // add this , otherwise main goroutine exit first
}
