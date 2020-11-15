package main

import "fmt"

func exercise1() {
	n := foo()
	n2, s := bar()
	fmt.Println(n, n2, s)
}

func foo() int {
	return 100
}

func bar() (int, string) {
	return 1, "Hello World"
}
