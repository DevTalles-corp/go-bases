package main

import (
	"fmt"
	"runtime"
	"time"
)

func count(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d -> iteración %d\n", id, i)
		runtime.Gosched() // Cambia a otra goroutine
	}
}

func main() {
	fmt.Println("Cores disponibles", runtime.NumCPU())
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	go count(1)
	go count(2)
	go count(3)
	go count(4)
	go count(5)

	time.Sleep(200 * time.Microsecond)
}
