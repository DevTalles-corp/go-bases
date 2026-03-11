package main

import "fmt"

func main() {
	ch := make(chan string) // Unbuffered

	go func() {
		fmt.Println("Goroutine: Voy a enviar...")
		ch <- "Mensaje listo"
		fmt.Println("Goroutine: Ya envié")
	}()

	fmt.Println("Main: Esperando a recibir...")
	msg := <-ch
	fmt.Printf("Main: recibí -> ", msg)
}
