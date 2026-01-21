package main

import "fmt"

const Pi float32 = 3.14159

func getTypedConstants() {
	var number float64 = float64(Pi)

	fmt.Println(number)
}
