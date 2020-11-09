package main

import (
	"fmt"
)

func exercise4() {
	y := 1994
	for {
		fmt.Println(y)
		if y == 2020 {
			break
		}
		y++
	}
}
