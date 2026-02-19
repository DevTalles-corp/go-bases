package main

import "fmt"

func main() {
	x := 10
	p := &x //Acceder a la dirección 0x0e3d

	fmt.Println("x = ", x)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	*p = 99
	fmt.Println("*p = ", *p)
	fmt.Println("x = ", x)
}
