package main

import "fmt"

func commaOkIdiom() {
	my_map := map[string]int{
		"Hola":  1,
		"Mundo": 2,
	}

	value, ok := my_map["Hola"]

	fmt.Println(value, ok)
}
