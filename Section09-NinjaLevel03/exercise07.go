package main

import (
	"fmt"
	"math/rand"
)

func exercise7() {
	num := rand.Int()
	if num == 0 {
		fmt.Println("Wow!")
	} else if num > 1 {
		fmt.Println("Yes!")
	} else {
		fmt.Println("Oh no!!!")
	}
}
