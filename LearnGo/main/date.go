package main
import (
	"fmt"
	"time"
)

func tryTime() {
	now := time.Now() // return a struct
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//formatted datetime
	dt := now.Format("2006-01-02 15:04:05") // the numbers in the string has to be fixed
	fmt.Printf("%T, %v", dt, dt)
}

