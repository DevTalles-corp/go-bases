package main

import "fmt"

func changeFirst(s []int) {
	s[0] = 7
}

func push(s []int) []int {
	s = append(s, 99)
	return s
}

func main() {
	a := []int{1, 2, 3}
	changeFirst(a)
	fmt.Println("a:", a)

	b := push(a)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
