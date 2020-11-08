package main

import "fmt"

func exercise4() {
	num := 123
	fmt.Printf("Decimal:%d\n", num)
	fmt.Printf("Binary:%b\n", num)
	fmt.Printf("Hex:%x\n", num)
	fmt.Println("---------------------")
	num = num << 1
	fmt.Printf("Decimal:%d\n", num)
	fmt.Printf("Binary:%b\n", num)
	fmt.Printf("Hex:%x\n", num)
}
