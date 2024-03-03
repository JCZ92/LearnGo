package main
import (
	"fmt"
	"strconv"
	"strings"
)

func tryString() {
	str := "hello world"
	fmt.Println(len(str)) // return bytes of the string
	// traverse the string by range
	for  i, v := range str {
		fmt.Printf("%d %c\n", i, v)
	}
	
	// traverse by slice
	ch := []rune(str)
	for i, v := range ch {
		fmt.Printf("%d %c\n", i, v)
	}

	// string to int
	n, _ := strconv.Atoi("123")
	fmt.Printf("type is %T, value is %v", n, n)

	// int to string
	n2 := 123
	str2 := strconv.Itoa(n2)
	fmt.Printf("type is %T, value is %v", str2, str2)

	// string to bool
	b, _ := strconv.ParseBool("true")
	fmt.Printf("type is %T, value is %v", b, b)

	// bool to string
	b2 := true
	str3 := strconv.FormatBool(b2)
	fmt.Printf("type is %T, value is %v", str3, str3)

	// string to float64
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("type is %T, value is %v", f, f)

	// float64 to string
	f2 := 3.14
	str4 := strconv.FormatFloat(f2, 'f', -1, 64)
	fmt.Printf("type is %T, value is %v", str4, str4)
	fmt.Println()
	fmt.Println(strings.Count("aaaaa", "aa"))// every char can be counted for only once


	// compare two strings ignoring case
	fmt.Println(strings.EqualFold("aaB", "AaB"))
	fmt.Println("aaB" == "AaB")


	// find the first index of the string
	fmt.Println(strings.Index("aaaaa", "aa"))
	fmt.Println(strings.LastIndex("aaaaa", "aa"))
	fmt.Println(strings.LastIndex("aaaaa", "ba")) // -1 means not found

	// replace the string
	fmt.Println(strings.Replace("aaaaa", "aa", "bb", -1))
	fmt.Println(strings.Replace("aaaaa", "aa", "bb", 1)) // only replace the first one

	// split the string, will return a slice
	splitted := strings.Split("a, b, c", ", ")
	strArr := [3]string{"a", "b", "c"}
	
	fmt.Printf("type is %T, value is %v", strArr, strArr) 
	fmt.Printf("type is %T, value is %v", splitted, splitted) 

	// toUpper case
	fmt.Println(strings.ToUpper("aBc"))

	// toLower case
	fmt.Println(strings.ToLower("aBc"))

	// trim the string
	fmt.Println(strings.Trim("  aBc  ", " "))
	fmt.Println(strings.TrimLeft("  aBc  ", " "))
	fmt.Println(strings.TrimRight("  aBc  ", " "))

	// join the string
	fmt.Println(strings.Join(splitted, "-"))
	// hasPrefix
	fmt.Println(strings.HasPrefix("aBc", "a"))
	// hasSuffix
	fmt.Println(strings.HasSuffix("aBc", "c"))
}