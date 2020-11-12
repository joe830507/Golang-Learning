package main

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

func main() {
	// myStruct()
	// embeddedStructs()
	anonymous()
}
