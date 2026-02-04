package main

import "fmt"

func forRangeLoop() {
	evenNumbers := []int{2, 4, 6, 8, 10, 12}

	for _, v := range evenNumbers {
		fmt.Println(v)
	}

	for i := range evenNumbers {
		fmt.Println(i)
	}
}
