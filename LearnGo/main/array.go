package main
import "fmt"

func arr() {
	// var array2 = [5]int{1,2,3,4,5} // you can initialize in this way
	// var array3 = []int{1, 3, 5}
	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = i
	}
	fmt.Println(array)
	fmt.Printf("array: %v\n", array)

	a := 0
	for _, v := range array { // traverse the array
		a += v
	}
	fmt.Println(a)
}

func compArray(arr [3]int) { // length is the attribute of the array, you cannot pass in an array of different length
	arr[0] = 1 // won't change the original array, passing by value
}

func compArray2(arr *[3]int) { // take a pointer to an array
	(*arr)[0] = 9
}

func twoDimenArray() {
	var array [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			array[i][j] = i + j
		}
	}
	
	fmt.Println(array)
	for _, v := range array {
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
	}
}



