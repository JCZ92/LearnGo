package main

import "fmt"

func control() {
	var a int = 10
	if (a == 10 || true != false) && 1 > 0 {
		fmt.Println("a is 10")
	}
	if b:= 1; b == 1 {
		fmt.Println("b is 1")
	}
	// fmt.Println("b is ", b)

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is not 1 or 2")
	}
}
