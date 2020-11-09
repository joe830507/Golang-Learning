package main

import "fmt"

func showSlicingASlice() {
	x := []int{1, 2, 3, 4, 5}
	// R : [1 2 3 4 5]
	fmt.Println(x[:])
	// R : [1 2 3 4 5]
	fmt.Println(x[0:])
	// It ends before index = 2.
	// R : [1 2]
	fmt.Println(x[:2])
	// It starts from index = 1, it ends before index = 2.
	// R : [2]
	fmt.Println(x[1:2])
}
