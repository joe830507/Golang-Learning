package main

import "fmt"

var str string = `This is an example sentence.`

func showStringType() {
	fmt.Printf("%T\n", str)
	b := []byte(str)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	for i := 0; i < len(str); i++ {
		fmt.Printf("hex:%#x, unicode:%#U, ASCII(base 10):%d\n", str[i], str[i], byte(str[i]))
	}
}
