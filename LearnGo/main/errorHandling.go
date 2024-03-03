package main
import "fmt"
import "errors"


func Divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero") // error
    }
    return x / y, nil
}

func handle1() {
	defer func() { // anonymous function
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}() // has to call it!!!
	result, err := Divide(10, 0)
	if err != nil {
		// fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println("Result:", result)
}

// defer and recover
func handle2() {
    defer func() {
		// recover() returns the value passed to panic() and stop the panic or nil if no panic occurred.
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    // Simulate a panic
    panic("something went wrong")
	fmt.Println("after panic") // this will not be executed due to unwinding the stack
}
