package main

import "fmt"

func showDefer() {
	defer shouldFirstPrint()
	shouldSecondPrint()
}

func shouldFirstPrint() {
	fmt.Println("First")
}

func shouldSecondPrint() {
	fmt.Println("Second")
}
