package encapsulation

type people struct {
	name string
	age  int
}

// factory function, like a constructor, note the naming convention
func NewPeople(name string, age int) *people {
	return &people{
		name: name,
		age:  age,
	}
}
// getter
func (p *people) GetName() string {
	return p.name
}

func (p *people) GetAge() int {
	return p.age
}

//setter
func (p *people) SetName(name string) {
	p.name = name
}

func (p *people) SetAge(age int) {
	p.age = age
}
