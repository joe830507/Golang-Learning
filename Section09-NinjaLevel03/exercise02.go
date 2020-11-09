package main

import (
	"fmt"
)

func exercise2() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%#U\n", i)
	}
}
