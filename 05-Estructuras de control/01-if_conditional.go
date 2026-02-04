package main

import "fmt"

func ifConditional() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}

	fmt.Println(x)
}
