package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Tarea cancelada: ", ctx.Err())
			return
		default:
			fmt.Println("Tarea trabajando...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	go task(ctx)

	time.Sleep(6 * time.Second)
	fmt.Println("Main terminó")
}
