package main

import "fmt"

const (
	cst1 = 1
	cst2 = 1232422.23123
	cst3 = "SSSSSS"
)

func showConstant() {
	fmt.Println(cst1)
	fmt.Println(cst2)
	fmt.Println(cst3)
	//It will cause an error.
	// cst1 = 2
	// fmt.Println(cst1)
}
