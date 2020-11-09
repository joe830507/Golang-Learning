package main

import "fmt"

//
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func showSwitch() {
	n := Tuesday
	switch n {
	case 1:
		fmt.Println("Today is Monday.")
	case 2:
		fmt.Println("Today is Tuesday.")
		fallthrough
	case 3:
		fmt.Println("Today is Wednesday.")
	case 4:
		fmt.Println("Today is Thursday.")
	case 5:
		fmt.Println("Today is Friday.")
	case 6:
		fmt.Println("Today is Saturday.")
	case 7:
		fmt.Println("Today is Sunday.")
	}
}
