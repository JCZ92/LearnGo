package main
import "fmt"

// camelCase starting with lowercase letter, package visiable

func cal (num1 int, num2 int) (int) { // note how to define the signature of the function
	return num1 + num2
}

// camelCase starting with uppercase letter,  exported, like public 

func Cal (num1 int, num2 int) int { // note how to define the signature of the function
	return num1 + num2
}
// func cal (num1 int) { function cannot be overload!!!

// }

// multiple return values
func cal2 (num1 int, num2 int) (int, int) { 
	return num1 + num2, num1 - num2
}
func cal3 (num1 int, num2 int)  { 
	tmp1, tmp2 := cal2(num1, num2)
	fmt.Println(tmp1, tmp2)
	tmp3, _ := cal2(num1, num2)
	fmt.Println(tmp3)
}
// multiple parameters
func cal4 (num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func functionAsVar() {
	ff := cal4;
	fmt.Printf("ff type is %T\n", ff);
	fmt.Println(ff(1, 2, 3, 4, 5));
}

func functionAsParameter(f func(int, int) int) { // looks complex
	fmt.Println(f(1, 2));
}

// define a type
type myFunction func(int, int) int

func functionAsParameter2(f myFunction) { // looks pretty
	fmt.Println(f(1, 2));
	fmt.Printf("f type is %T\n", f);
}


// casting and type
type myInt int;

var intA myInt = 1;
var intB int = 2;

func cast() {
	// myInt != int in compilation
	fmt.Println(int(intA) + intB); // have to cast either one
	fmt.Println(intA + myInt(intB));  // the mirror
}

// name the return value
func cal5 (num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

// anonymous function
func anonymousFunction() {
	f := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println(f(1, 2));

	func (num1 int) {
		fmt.Println(num1);
	}(6)
}
