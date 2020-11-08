package main

import "fmt"

var complx complex64

func showNumericTypes() {
	num1 := 42
	num2 := 42.33334
	fmt.Printf("%T\n", num1)
	fmt.Printf("%T\n", num2)
	fmt.Printf("Values:%v,Type:%T\n", complx, complx)
	complx = 100.34234
	fmt.Printf("Values:%v,Type:%T\n", complx, complx)
}
