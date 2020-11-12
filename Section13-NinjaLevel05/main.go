package main

type person struct {
	first          string
	last           string
	iceCreamFlavor string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	exercise4()
}
