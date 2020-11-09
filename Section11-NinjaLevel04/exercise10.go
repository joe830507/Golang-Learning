package main

import (
	"fmt"
	"strings"
)

func exercise10() {
	m := map[string]([]string){
		"Joe":   []string{"Tofu", "Bubble Drink"},
		"John":  []string{"Video Games", "Paintings"},
		"Larry": []string{"Soccer", "Swimming"},
	}
	m["Johnson"] = []string{"Basketball", "Football"}
	for name, v := range m {
		str := ""
		str += strings.Join(v, ", ")
		fmt.Println(name, str)
	}
	fmt.Println("---------------------")
	delete(m, "Johnson")
	for name, v := range m {
		str := ""
		str += strings.Join(v, ", ")
		fmt.Println(name, str)
	}
}
