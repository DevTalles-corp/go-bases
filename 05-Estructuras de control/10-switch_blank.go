package main

import "fmt"

func switchBlank() {
	a := 3

	switch {
	case a == 2:
		fmt.Println("a es 2 dos dos")
	case a == 3:
		fmt.Println("a es 3 tres tres")
	default:
		fmt.Println("a es ", a)
	}
}
