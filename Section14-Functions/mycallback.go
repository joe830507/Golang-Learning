package main

import "fmt"

func showCallback() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sumValues(nums...)
	fmt.Println(total)
	r := even(sumValues, nums...)
	fmt.Println(r)
}

//define a func
func sumValues(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//passing a function "f func(xi ...int) int"
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
