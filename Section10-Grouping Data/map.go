package main

import "fmt"

func showMap() {
	// key : string, value : int
	m := map[string]int{
		"First":  1,
		"Second": 2,
		"Third":  3,
	}
	fmt.Println(m)
	fmt.Println(m["First"])
	fmt.Println(m["Second"])
	for i, v := range m {
		fmt.Println(i, v)
	}

	v1, ok1 := m["First"]
	fmt.Println(v1)
	fmt.Println(ok1)

	if v1, ok1 := m["Second"]; ok1 {
		fmt.Println(v1)
	}

	v2, ok2 := m["Four"]
	fmt.Println(v2)
	fmt.Println(ok2)

	if v2, ok2 := m["Four"]; ok2 {
		fmt.Println(v2)
	} else {
		fmt.Println("Nothing.")
	}
	// additional
	fmt.Println(len(m))
}
