package main

import "fmt"

func showReturningFunc() {
	sf := returnFunc()
	//It's a func() int
	fmt.Printf("%T\n", sf)
	fmt.Println(sf())
	fmt.Println(returnFunc()())
}

func returnFunc() func() int {
	return func() int {
		return 42
	}
}
