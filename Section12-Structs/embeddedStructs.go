package main

import "fmt"

func embeddedStructs() {
	p1 := secretAgent{
		person: person{
			first: "Joe",
			last:  "Lin",
			age:   25,
		},
		first: "Something to test duplicate keys",
		ltk:   true,
	}
	p2 := person{
		first: "Lulu",
		last:  "Lala",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.person.first, p1.last, p1.age, p1.first, p1.ltk)
	fmt.Println(p2.first, p2.last)
}
