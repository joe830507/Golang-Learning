package main

import "fmt"

func anonymous() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Joe",
		last:  "Lin",
		age:   25,
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
