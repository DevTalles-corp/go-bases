package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(7 * time.Second)
		ch1 <- "Respuesta de channel 1"
	}()

	go func() {
		time.Sleep(7 * time.Second)
		ch2 <- "Respuesta de channel 2"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Recibí de ch1: ", msg)
		case msg := <-ch2:
			fmt.Println("Recibí de ch2: ", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: ningun channel respondió a tiempo")
		}
	}
}
