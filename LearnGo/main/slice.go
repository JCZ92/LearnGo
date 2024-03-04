package main
import "fmt"

func slice() {
	// the underlying data structure of a slice is an array, with a length and capacity
	var intarr [5]int = [5]int{1, 2, 3, 4, 5}
	var slice []int = intarr[1:3]// declaration 1
	// slice2 := intarr[1:3] 
	// slice := make([]int, 3, 5) // declaration 2, type, len, cap
	// slice := []int{1, 2, 3} // declaration 3, type, values
	// slice := []int{1, 2, 3}
	
	fmt.Println(slice)
	fmt.Println("slice length is", len(slice))
	//capacity of a slice is the length of the underlying array
	fmt.Println("slice capacity is", cap(slice))
	manSlice(slice) // passing copy of slice header(reference to array, len, cap), but can modify the underlying array
	fmt.Println(intarr)
	fmt.Println(slice) // slice is changed in the function, but not the original slice
	
	// Original slice of slices
	original := [][]int{{1, 2}, {3, 4}}

	// Create a new slice and copy the original slice to it
	copied := make([][]int, len(original))
	copy(copied, original) // a shallow copy

	// Modify an element in the copied slice
	copied[0][0] = 100

	// Print both slices
	fmt.Println("Original:", original) // Output: [[100 2] [3 4]]
	fmt.Println("Copied:", copied)     // Output: [[100 2] [3 4]]

	// delete a slice element, slice2[2] = slice2[3] and slice2[3:] = slice2[4:]
	slice2 := []int{1, 2, 3, 4, 5}
	slice2 = append(slice2[:2], slice2[3:]...) // Delete element at index 2, this is very verbose
	fmt.Println(slice2)
}

func manSlice(slice []int) { // passing by slice header
	slice[0] = 10 // will change the underlying array, but won't change the passed-in slice when the function returns
	fmt.Println(slice)
	//fmt.Println(slice[3])// this is not allowed, indexing out of range
	// but this is okay, slicing beyond the capacity, whether to use this style depends on your choice
	fmt.Println(slice[0:3])
	slice = append(slice, 6, 7, 8, 9, 10) // expand the length, but won't change the passed-in slice when the function returns
	fmt.Println(slice)
}
