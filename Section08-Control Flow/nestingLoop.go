package main

import "fmt"

func showNestingLoop() {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println(i * j)
		}
	}
}
