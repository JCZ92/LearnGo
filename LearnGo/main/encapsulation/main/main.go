package main

import (
	"LearnGo/LearnGo/main/encapsulation"
	"fmt"
)

func main() {
	p := encapsulation.NewPeople("Angla", 10)
	fmt.Println(*p)
	p.SetAge(20)
	fmt.Println(*p)
}
