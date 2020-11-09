package main

import "fmt"

func showAppendSlice() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6, 7, 8)
	// R : [1 2 3 4 5 6 7 8]
	fmt.Println(x)
	y := []int{9, 10, 11, 12}
	x = append(x, y...)
	// R : [1 2 3 4 5 6 7 8 9 10 11 12]
	fmt.Println(x)
}
