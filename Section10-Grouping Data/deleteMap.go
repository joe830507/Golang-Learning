package main

import "fmt"

func showDeleteMap() {
	// key : string, value : int
	m := map[string]int{
		"First":  1,
		"Second": 2,
		"Third":  3,
	}
	m["Four"] = 4
	fmt.Println(m)

	delete(m, "Four")
	fmt.Println(m)
	delete(m, "Five")
	fmt.Println(m)
	if v, ok := m["Four"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Nothing.")
	}
}
