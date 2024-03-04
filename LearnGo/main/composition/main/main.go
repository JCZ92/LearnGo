package main

import (
	"LearnGo/LearnGo/main/composition"
	"fmt"
)


func main() {
	dog := composition.NewDog()

	dog.Speak() // dog can call method of animal
	dog.Yell()
	dog.Eat()
	fmt.Println(dog.Name) // dog can access field of animal
	// fmt.Println(dog.int) // dog can't access its anonymous field from other package
}

