package main

import (
	"flag"
	"fmt"

	"github.com/ricardocuellar/greetctl/internal/logx"
	gr "github.com/ricardocuellar/greetctl/pkg/greet"
)

func main() {
	name := flag.String("name", "mundo", "Nombre a saludar")
	flag.Parse()

	msg := gr.Hello(*name)
	fmt.Println("Proyecto listo desde CMD!")
	fmt.Println(msg)

	logx.Info("FIN todo OK")
}
