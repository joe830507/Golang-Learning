package main

import "fmt"

var tempBoolType bool

func showBoolType() {
	fmt.Println(tempBoolType)
	tempBoolType = true
	fmt.Println(tempBoolType)
	x := 7
	y := 42
	fmt.Println(x == y)
	y = 7
	fmt.Println(x == y)
}
