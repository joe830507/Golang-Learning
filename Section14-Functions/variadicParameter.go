package main

import "fmt"

func showVariadicParameter() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := caculateTotal(nums...)
	fmt.Println(sum)
}

func caculateTotal(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
