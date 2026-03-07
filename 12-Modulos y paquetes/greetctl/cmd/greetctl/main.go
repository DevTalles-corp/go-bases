package main

import (
	"fmt"

	"github.com/ricardocuellar/greetctl/internal/logx"
	gr "github.com/ricardocuellar/greetctl/pkg/greet"
)

func main() {
	msg := gr.Hello("      Fernando       ")
	fmt.Println("Proyecto listo desde CMD!")
	fmt.Println(msg)

	logx.Info("FIN todo OK")
}
