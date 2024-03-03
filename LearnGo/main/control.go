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

	switch a % 2 {
	case 1:
		fmt.Println("a is 1") // no break is required
	case 2, 3, 4: // can have multiple values
		fmt.Println("a is 2")
	default: // default is not mandatory
		fmt.Println("a is not 1 or 2")
	}

	switch { // for complex conditions, w/o specifying the variable
		case a > 10:
			fmt.Println("a > 10") // no break is required
		case a < 10:
			fmt.Println("a < 10")
		default:
			fmt.Println("a = 10")
	}

	for i:= 1 ; i < 5; i++ { // cannot use var here
		fmt.Println(i)
		var j int = i * 10
		fmt.Println(j)
	}
	// for {
		// this is infinite loop
	// }


}	
