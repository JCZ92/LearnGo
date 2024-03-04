package main

import "fmt"

func closure() {
    // Define a function that returns another function
    // The returned function captures the variable 'x' from the outer scope
    add := func() func(int) int {
		sum := 0
        return func(y int) int {
			sum += y
            return sum
        }
    }
	f := add() // closure is the function and its context(in this case, its context is variable sum)
	fmt.Println(f(1))
	fmt.Println(f(1))
	// closure will consume memory heavily


	// note that below has different effect, we are creating new function
	fmt.Println(add()(1))
	fmt.Println(add()(1))
}
