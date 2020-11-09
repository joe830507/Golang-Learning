package main

import (
	"fmt"
	"strings"
)

func exercise8() {
	m := map[string]([]string){
		"Joe":   []string{"Tofu", "Bubble Drink"},
		"John":  []string{"Video Games", "Paintings"},
		"Larry": []string{"Soccer", "Swimming"},
	}
	for name, v := range m {
		str := "Favorite Things: "
		str += strings.Join(v, ", ")
		fmt.Println(name, str)
	}
}
