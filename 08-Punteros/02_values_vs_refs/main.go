package main

import "fmt"

func incrementValue(number int) {
	number++
}

func mutateFirst(my_slice []int) {
	my_slice[0] = 999
}

func mutateSecond(my_slice []int) {
	my_slice[1] = 1995
}

func push(s []int) []int {
	s = append(s, 42)
	return s
}

func main() {
	x := 10
	incrementValue(x) //11?
	fmt.Println("x: ", x)

	a := []int{1, 2, 3}
	mutateFirst(a)
	fmt.Println("a: ", a)

	b := push(a)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	mutateSecond(b)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

}
