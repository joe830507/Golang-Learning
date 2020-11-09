package main

import "fmt"

func showMake() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x)) //10
	fmt.Println(cap(x)) //12

	// If we put a value into the slice and its size is over the capacity,
	// the capacity of the slice will be bigger 2 times as its new capacity.
	// The current length is 13, so the new capacity is 12 * 2 = 24.
	x = append(x, 3423) //length:11, capacity:12
	x = append(x, 3423) //length:12, capacity:12
	x = append(x, 3423) //length:13, capacity:24
	fmt.Println(x)
	fmt.Println(len(x)) //13
	fmt.Println(cap(x)) //24

}
