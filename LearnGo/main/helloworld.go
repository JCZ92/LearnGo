package main
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello, World") // no need semicolon

	// var age int = 10 
	// fmt.Println("age is" , age, "very good")
	// var age2 int
	// fmt.Println("age2 is", age2)
	// var age3 = "20.3" // without type declaration
	// fmt.Println("age3 is", age3)
	// age4:= 9 // w/o var
	// fmt.Println(age4)

	// var v1, v2, v3 int
	// fmt.Println(v1, v2, v3)
	// var v4, v5, v6 = 1, '2', 3.9
	// fmt.Println(v4, v5, v6)
	// v7, v8:= 10, 11
	// fmt.Println(v7, v8)
	// var (
	// 	v9 = 12
	// 	v10 = 13
	// )
	// fmt.Println(v9, v10)


	var size int16 = 127;
	// fmt.Println(size + 1)
	fmt.Println(unsafe.Sizeof(size))// output bytes of variable
	// var str string = "ab" + "cd"
	// fmt.Println(str) 
	// fmt.Println(10 / 2.4)

	// a % b = a - (a/b)*b
	// fmt.Println(7 % 4)
	// fmt.Println(-7 % 4)
	// fmt.Println(7 % -4)
	// fmt.Println(-7 % -4)
	// fmt.Printf("type is %T", true)
	// var ptr *int = &age;
	// fmt.Println(ptr)
	// fmt.Println(*ptr)	
	// *ptr = 20;
	// fmt.Println(age)
	// control()
	// fmt.Println(cal(1, 2))
	// fmt.Println(cal2(1, 2))
	// functionAsVar()
	// functionAsParameter(cal)
	// functionAsParameter2(cal)
	// cast()
	// handle1()
	arr()
	var array3 = [...]int{1, 3, 5} // array
	compArray(array3)
	fmt.Println(array3)
	compArray2(&array3)
	fmt.Println(array3)
	twoDimenArray()
	slice()

}
