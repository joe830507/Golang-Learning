package main

import "fmt"

func exercise6() {
	names := make([]string, 5, 10)
	y := []string{
		"James",
		"Joe",
		"John",
		"Ruru",
		"Larry",
	}
	for i, v := range y {
		names[i] = v
	}
	for i, v := range names {
		fmt.Println(i, v)
	}
	fmt.Println(len(names))
	fmt.Println(cap(names))
}
