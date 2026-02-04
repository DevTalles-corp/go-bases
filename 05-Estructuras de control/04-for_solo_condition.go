package main

import "fmt"

func forSoloCondition() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}
