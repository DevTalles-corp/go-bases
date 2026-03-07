package main

import (
	"fmt"

	"github.com/ricardocuellar/greetctl/pkg/greet"
)

func main() {
	msg := greet.Hello("Fernando")
	fmt.Println("Proyecto listo desde CMD!")
	fmt.Println(msg)
}
