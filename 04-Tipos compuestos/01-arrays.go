package main

import "fmt"

func arrays() {
	// var number_list [3]int
	//Array literal
	// var number_list = [3]int{10, 20, 30}
	//Array literal simplificado
	var number_list = [...]int{10, 20, 30, 40}
	fmt.Println(number_list)

	number_list[0] = 50
	fmt.Println(number_list[2])

	fmt.Println(len(number_list))
}
