package main

import "fmt"

const (
	aaa = iota
	bbb
	ccc
)

func showIota() {
	fmt.Printf("Type:%T, Value:%v\n", aaa, aaa)
	fmt.Printf("Type:%T, Value:%v\n", bbb, bbb)
	fmt.Printf("Type:%T, Value:%v\n", ccc, ccc)
}
