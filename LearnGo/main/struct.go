package main
import	"fmt"

// go has no class
type Person struct { // PascalCase means exported types
	Name string // field declaration, note that we cannot set default values
	address string
	Age  int
	salary int // camelCase means private access to the package

}

type Student struct {
	Name string // field declaration
	address string
	Age  int
	salary int 
	// school string
}

func tryStruct() {
	var john Person = Person {Name: "John"}
	// var john Person = Person {"John", "main st", 10, 1000} // has to initialize all fields by this way
	fmt.Println(john)
	john.Name = "John"
	john.address = "123 Main St"
	john.Age = 30
	john.salary = 1000
	fmt.Println(john)

	var may *Person = new(Person)
	// var may *Person = &Person{}
	may.Name = "May"
	may.address = "456 Main St"
	may.Age = 20
	may.salary = 900
	fmt.Println(may)
	fmt.Println(maxx)


	// conversion of struct instance
	var jane Student = Student(john) // copy all fields from john to jane when 2 structs have exact same fields
	fmt.Println(jane)
}

func (p Person) tryMethod() {
	fmt.Println("Hello,", p.Name)
	p.Name = "John" // pass by value and won't change the outside object
}
func (p *Person) tryMethodByPointer() {
	fmt.Println("Hello,", p.Name)
	p.Name = "John" // will modify the outside object
}

type myIntt int

func (i myIntt) double() int { // method can be defined on customized types
	return int(i * 2)
}