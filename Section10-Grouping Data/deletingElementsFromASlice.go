package main

import "fmt"

func showDeletingElements() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 99, 88, 77, 66, 55)
	y := []int{111, 222, 333, 444}
	x = append(x, y...)
	// [1 2 3 4 5 99 88 77 66 55 111 222 333 444]
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	// [1 2 5 99 88 77 66 55 111 222 333 444]
	fmt.Println(x)
}
