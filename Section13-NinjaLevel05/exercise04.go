package main

import "fmt"

func exercise4() {
	v := struct {
		doors int
		color string
	}{
		doors: 4,
		color: "red",
	}
	fmt.Println(v)
}
