package main

import "fmt"

func sliceFunctions() {
	// var my_slice []int
	// fmt.Println(my_slice == nil)

	// my_slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(my_slice)
	// //len
	// fmt.Println(len(my_slice))

	// //append
	// my_slice = append(my_slice, 10, 12, 30, 12, 40, 50, 12, 3, 12, 35, 12, 13)
	// fmt.Println(my_slice)

	// //Capacity
	// fmt.Println(my_slice, len(my_slice), cap(my_slice))
	// fmt.Println(len(my_slice) == cap(my_slice))

	// make_slice := make([]int, 5)
	make_slice := make([]int, 0, 10)
	fmt.Println(make_slice)

	make_slice = append(make_slice, 10, 20, 30, 40, 10, 20, 30, 40, 10, 20, 30, 40)
	fmt.Println(make_slice)
	fmt.Println(len(make_slice), cap(make_slice))

	// Vaciar slice
	clear(make_slice)
	fmt.Println(make_slice)
	fmt.Println(len(make_slice), cap(make_slice))

}
