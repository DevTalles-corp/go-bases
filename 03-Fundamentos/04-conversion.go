package main

import "fmt"

func getConversion() {
	var number1 int = 10
	var number2 float64 = 3.5

	total := float64(number1) + number2
	fmt.Println(total)
}
