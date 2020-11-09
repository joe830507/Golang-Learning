package main

import "fmt"

func showIf() {
	if 23132423%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 10; num < 9 {
		fmt.Println("error")
	} else if num == 10 {
		fmt.Println("success")
	} else {
		fmt.Println("error")
	}
}
