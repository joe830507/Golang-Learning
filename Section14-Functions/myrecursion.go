package main

import "fmt"

func showRecursion() {
	n := factorial(4)
	fmt.Println(n)
	fmt.Println(fibonacci(10))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
