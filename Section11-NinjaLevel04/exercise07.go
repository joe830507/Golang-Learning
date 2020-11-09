package main

import "fmt"

func exercise7() {
	str1 := []string{"James", "Bond", "Shaken, not stirred"}
	str2 := []string{"Miss", "Moneypenny", "Hellooooooo, James"}
	str3 := [][]string{str1, str2}
	for _, v1 := range str3 {
		str := ""
		for _, v2 := range v1 {
			str += v2 + " "
		}
		fmt.Println(str)
	}
}
