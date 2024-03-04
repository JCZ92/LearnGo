package composition
import (
	"fmt"
)

type Dog struct {
	Animal // anonymous struct
	int
}

func (dog *Dog) Eat() {
	fmt.Println(dog.Name, "is eating")
}
func (dog *Dog) Yell() {
	fmt.Println(dog.Name, "Woof", dog.int)
}

func NewDog() (*Dog) {
	return &Dog{Animal{Name: "Dog"}, 99}
}

