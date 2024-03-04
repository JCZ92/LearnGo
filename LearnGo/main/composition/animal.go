package composition

import "fmt"

type Animal struct {
	Name string
}

func (animal *Animal) Yell() {
	fmt.Println(animal.Name, "yelling")
}

func (animal *Animal) Speak() {
	fmt.Println("My name is", animal.Name)
}
