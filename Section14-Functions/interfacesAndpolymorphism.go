package main

import "fmt"

func (p person) speak() {
	fmt.Println("I'm", p.first, p.last)
}

// a value which has speak() method, it's also the type human
// a value can be more than one type.
type human interface {
	speak()
}

func saySomething(h human) {
	//human might be the type person or secretAgent
	switch h.(type) {
	case person:
		fmt.Println("I'm an ordinary person.--------------", h.(person).first)
	case secretAgent:
		fmt.Println("I'm a secretAgent.-----------------", h.(secretAgent).first)
	}
	fmt.Println("I called human")
}

//according to the type human interface
//p1 can not only be secretAgent, but it can be the type human.
func myInterfaces() {
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
		first: "Marry",
		last:  "Li",
		age:   28,
	}
	p1.speak()
	fmt.Printf("%T\n", p1)
	saySomething(p1)
	saySomething(p2)
}
