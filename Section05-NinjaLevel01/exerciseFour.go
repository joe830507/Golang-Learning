package main

import "fmt"

type myOwnType int

var xx myOwnType

func exerciseFour() {
	fmt.Printf("%v\n", xx)
	fmt.Printf("%T\n", xx)
	xx = 42
	fmt.Println(xx)
}
