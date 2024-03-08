package main
import (
	"LearnGo/LearnGo/main/interface" // path where the package is defined
	"fmt"
)

func welcome(s interface1.SaySth) { 
	// the key here is that taking a interface, wanting s implements interface SaySth
	// any type that implements interface SaySth can be passed to this function
	s.SayHello(1) // note that we are passing by pointer, may have side effects!!!!
	
	// type assertion
	gm, ok  := s.(interface1.GoodMan)
	if ok {
		gm.SayNo()
	}

	// type switch is also useful
	switch s.(type) {
	case interface1.GoodMan:
		fmt.Println("GoodMan!")
	case interface1.BadMan:
		fmt.Println("BadMan")
	default:
		fmt.Println("Unknown")
	}
}


func main() {	


	gm := interface1.GoodMan{Num: 8}
	// gm2 := &interface1.GoodMan{Num: 8}
	// bm := interface1.BadMan{}
	fmt.Println(gm.Num)
	welcome(gm)
	// welcome(gm2)
	// welcome(bm) // bm does not implement interface SaySth (missing SayNo method), a loose contract

	// use interface to initialize a variable
	// var s interface1.SaySth = interface1.GoodMan{Num: 8}
	// welcome(s)

	// var emptyInterface interface{} = s; // an empty interface can take any type
	// fmt.Println(emptyInterface)
	var val interface{} // really interesting
	val = 42
	val = "hello"
	val = true
	val = []int{1,2,3}
	fmt.Println(val)

	//polymorphism array

	// var arrPM [3]interface1.SaySth
	// arrPM[0] = interface1.GoodMan{Num: 8}
	// arrPM[1] = interface1.GoodMan{Num: 0}
	


}

