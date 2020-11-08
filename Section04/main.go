package main

import "fmt"

func main() {
	assignValueToVariable()
	declaredMyType()
}

var x string

func assignValueToVariable() {
	x = "Hello World"
	y := "Hello World"
	fmt.Println(x)
	fmt.Println(y)
}

type myType int

var z myType

func declaredMyType() {
	fmt.Printf("Value:%v,Type:%T\n", z, z)
	z = 100
	fmt.Printf("Value:%v,Type:%T\n", z, z)
	s := int(z)
	fmt.Printf("Value:%v,Type:%T\n", s, s)
}
