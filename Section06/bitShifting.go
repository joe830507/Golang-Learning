package main

import "fmt"

const (
	iii = iota             // 0
	kb  = 1 << (iota * 10) // 1 << ( 1 * 10 )
	mb  = 1 << (iota * 10) // 1 << ( 2 * 10 )
	gb  = 1 << (iota * 10) // 1 << ( 3 * 10 )
)

func showBitShifting() {
	// 2^2
	x := 4
	fmt.Printf("%v \t %b\n", x, x)
	// 2^3
	y := x << 1
	fmt.Printf("%v \t %b\n", y, y)
	// 2^5
	z := y << 2
	fmt.Printf("%v \t %b\n", z, z)
	aa := 15
	fmt.Printf("%v \t %b\n", aa, aa)
	bb := aa << 1
	fmt.Printf("%v \t %b\n", bb, bb)
	one := 1
	fmt.Printf("%v \t %b\n", one, one)
	zero := one >> 2
	fmt.Printf("%v \t %b\n", zero, zero)

	fmt.Println("------------------------")
	fmt.Println(iii)
	fmt.Println(1 << 10)
	fmt.Println(kb)
	fmt.Println(1 << 20)
	fmt.Println(mb)
	fmt.Println(1 << 30)
	fmt.Println(gb)
}
