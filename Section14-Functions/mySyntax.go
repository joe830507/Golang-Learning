package main

import "fmt"

func showSyntax() {
	foo()
	bar("James")
	s := woo("Susan")
	fmt.Println(s)
	a, b := mouse("Tina", "Larry")
	fmt.Println(a)
	fmt.Println(b)
}

func foo() {
	fmt.Println("Hello world.")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprintln("hello,", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ",", ln, ` says "hello"`)
	b := true
	return a, b
}
