package main

import "fmt"

func exercise2() {
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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Println(v.first, v.last, v.iceCreamFlavor)
	}
	fmt.Println(m["Bond"].first)
	fmt.Println(m["Bond"].last)
	fmt.Println(m["Bond"].iceCreamFlavor)
	fmt.Println(m["Kimberley"].first)
	fmt.Println(m["Kimberley"].last)
	fmt.Println(m["Kimberley"].iceCreamFlavor)
}
