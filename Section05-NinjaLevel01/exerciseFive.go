package main

import "fmt"

var xxx myOwnType

var yy int

func exerciseFive() {
	fmt.Printf("Value:%v\n", xxx)
	fmt.Printf("Type:%T\n", xxx)
	xxx = 100
	fmt.Printf("Value:%v\n", xxx)
	yy = int(xxx)
	fmt.Printf("Value:%v\n", yy)
}
