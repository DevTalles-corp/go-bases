package main

import (
	"fmt"
	"slices"
)

func slicesExample() {
	// Slice literal
	// var my_slice = []int{1, 2, 3}
	// fmt.Println(my_slice)

	// my_slice[1] = 23
	// fmt.Println(my_slice)
	// fmt.Println(my_slice[1])

	// Slice vacio
	// var zero_slice []int
	// var other_zero_slice []int
	// fmt.Println(zero_slice == nil)

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	// s := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	//fmt.Println(slices.Equal(x, s)) //NO Compila
}
