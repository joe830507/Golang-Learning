package main

import "fmt"

//the container of variables
//some variables are limited in the scope.
func showClosure() {

	x := 111
	fmt.Println(x)
	{
		y := 222
		fmt.Println(y)
	}
	// it can't be executed because y is inside the bracket.
	// y can't be read by outside codes.
	// fmt.Println(y)

	// x still exist because the func(includes integer x)
	// is assigned to the variable 'increment'
	// if we want x to be changed, execute callback func
	// the func has access to change x.
	incrementA := incrementor()
	// incrementB has the scope which is different from incrementA.
	incrementB := incrementor()
	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementB())
	fmt.Println(incrementB())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
