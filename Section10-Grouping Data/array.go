package main

import "fmt"

func showArray() {
	var a [5]int
	fmt.Println(a)
	a[3] = 42
	fmt.Println(a)
	fmt.Println(len(a))
}
