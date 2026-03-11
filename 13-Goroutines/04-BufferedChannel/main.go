package main

import "fmt"

func main() {
	ch := make(chan int, 3) // Buffer es de capacidad 3

	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40

	fmt.Println("Envié 3 valores sin bloquearme")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
