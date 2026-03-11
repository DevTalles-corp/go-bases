package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Println("Hola, ", name)
}

func main() {
	go greet("Ricardo")
	go greet("Ana")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("main terminó")

}
