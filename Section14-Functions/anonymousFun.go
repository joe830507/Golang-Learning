package main

import "fmt"

func showAnonymousFun() {
	func() {
		fmt.Println("This is a anonymous func.")
	}()
	func(x int) {
		fmt.Println("Show number", x, ".")
	}(100)
}
