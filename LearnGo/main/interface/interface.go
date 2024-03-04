package interface1

type SaySth interface {
	SayHello(num int)
}

type EnhancedInterface interface {
	SaySth // reuse the SaySth interface
	SayGoodbye()
}
