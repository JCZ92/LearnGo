package main
import "fmt"

func mapType() {
	// declaration
	var m1 map[int]string // declration won't allocate memory
	m1 = make(map[int]string, 10) // initialize by make
	// m1 := make(map[int]string) // initialize by make
	m1[1] = "a" // add key-value
	m1[2] = "b"
	m1[3] = "c"
	fmt.Println(m1)
	// initialze a map
	m2 := map[int]string{1: "a", 2: "b", 3: "c",}
	fmt.Println(m2)

	m1[1] = "d"// update value
	fmt.Println(m1)

	delete(m1, 1)// delete key - value
	// there is no built-in way to delete all keys
	// you can iterate over the map and delete each key

	fmt.Println(m1)
	v2 := m1[2]
	fmt.Println(v2)

	// check length of map
	fmt.Println(len(m1))

	// iterate over the map
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// check if key exists
	if _, exists := m1[1]; exists {
		fmt.Println("key exists")
	}

	originalMap := map[string][]int{"a": []int{1, 2}, "b": []int{31, 32}, "c": []int{11, 21}}

    // Create a copy of the original map, note that it is a shallow copy
    copiedMap := make(map[string][]int)
    for key, value := range originalMap {
        copiedMap[key] = value
    }
	copiedMap["b"][0] = 33 // update the value of the key "b" in the copied map

    fmt.Println("Original map:", originalMap)
    fmt.Println("Copied map:", copiedMap)
}