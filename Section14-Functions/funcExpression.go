package main

import "fmt"

func showFuncExpression() {
	f := func() {
		fmt.Println("This is a anonymous func.")
	}
	f()
	f2 := func(x int) {
		fmt.Println("Show number", x, ".")
	}
	f2(100)
}
