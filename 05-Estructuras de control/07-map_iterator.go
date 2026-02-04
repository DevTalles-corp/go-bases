package main

import "fmt"

func mapIterator() {
	my_map := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		println("Loop: ", i)
		for k, v := range my_map {
			fmt.Println(k, v)
		}
	}
}
