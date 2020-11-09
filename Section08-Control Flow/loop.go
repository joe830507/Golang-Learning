package main

import (
	"fmt"
)

func showFor() {
	num := 0
	// while(num < 10)
	for num < 10 {
		fmt.Println(num)
		num++
	}

	a := 10

	// while(true)
	for {
		a++
		fmt.Println("break!")
		if a == 20 {
			break
		}
	}

	a = 10

	for {
		if a == 15 {
			a++
			continue
		} else if a == 20 {
			break
		} else {
			a++
			fmt.Println(a)
		}
	}
}
