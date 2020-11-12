package main

import "fmt"

func myStruct() {
	p1 := person{
		first: "Joe",
		last:  "Lin",
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last)
}
