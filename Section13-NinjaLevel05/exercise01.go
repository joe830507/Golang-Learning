package main

import "fmt"

func exercise1() {
	p1 := person{
		first:          "James",
		last:           "Bond",
		iceCreamFlavor: "cholocate",
	}
	p2 := person{
		first:          "Marry",
		last:           "Kimberley",
		iceCreamFlavor: "Herb",
	}
	ps := []person{p1, p2}
	for _, v := range ps {
		fmt.Println(v.first, v.last, v.iceCreamFlavor)
	}
}
