package main

import (
	"fmt"
	"sync"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d trabajando...\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}

	wg.Wait()
	fmt.Println("Todos los workers terminaron")
}
