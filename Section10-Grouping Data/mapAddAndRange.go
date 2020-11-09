package main

import "fmt"

func showMapAddAndRange() {
	// key : string, value : int
	m := map[string]int{
		"First":  1,
		"Second": 2,
		"Third":  3,
	}
	m["Four"] = 4
	fmt.Println(m)
	for i, v := range m {
		fmt.Println(i, v)
	}

}
