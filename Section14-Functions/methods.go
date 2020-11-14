package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

//func (r receiver) identifier(parameters) (return(s)) { code }

//attach speak() to the type 'secretAgent'
//It's kind of like C# extensions.
func (s secretAgent) speak() {
	fmt.Println("I'm", s.person.first, s.person.last)
}

func myMethods() {
	p1 := secretAgent{
		person: person{
			first: "Joe",
			last:  "Lin",
			age:   25,
		},
		first: "Something to test duplicate keys",
		ltk:   true,
	}
	p2 := secretAgent{
		person: person{
			first: "Marry",
			last:  "Li",
			age:   27,
		},
		first: "Something to test duplicate keys",
		ltk:   true,
	}
	p1.speak()
	p2.speak()
}
