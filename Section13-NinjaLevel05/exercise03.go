package main

import "fmt"

func exercise3() {
	v := vehicle{
		doors: 4,
		color: "red",
	}
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	fmt.Println(v)
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(v.doors, v.color)
	fmt.Println(t.doors, t.color, t.fourWheel)
	fmt.Println(s.doors, s.color, s.luxury)
}
