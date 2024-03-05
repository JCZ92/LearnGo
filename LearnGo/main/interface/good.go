package interface1

type GoodMan struct {
	Num int
}

func (g GoodMan) SayHello(num int) {
	g.Num = num
	println("A good man say hello", g.Num)
}

// func (g *GoodMan) SayHello(num int) { // only to use pointer receiver to update field
// 	g.Num = num
// 	println("A good man say hello", g.Num)
// }

func (g GoodMan) SayNo() {
	println("A good man say sorry")
}
type BadMan struct {

}

func (b BadMan) SayHello(num int) {
	println("Yoo?")
}
