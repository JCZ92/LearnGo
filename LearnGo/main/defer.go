package main

import "fmt"

// defer statement will be pushed into a stack
// when the function returns, the stack will be executed in [reverse] order
// variable value of defer statements won't be changed by later operations on the same variable
// can be used for closing resources

func tryDefer(num1 int, num2 int) {
	defer fmt.Println(num1) 
	defer fmt.Println(num2)

	num1 *= 2 // will not change value in previous defer
	num2 *= 2
	sum := num1 + num2
	fmt.Println(sum)
	return 
}