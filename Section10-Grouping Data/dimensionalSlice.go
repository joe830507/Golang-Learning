package main

import "fmt"

func showDimensionalSlice() {
	x1 := []int{1, 2, 3, 4, 5}
	x2 := []int{6, 7, 8, 9, 10}
	x3 := [][]int{x1, x2}
	fmt.Println(x3)
	// R: [[1 2 3 4 5] [6 7 8 9 10]]
}
