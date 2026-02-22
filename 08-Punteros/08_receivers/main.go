package main

import "fmt"

type Counter struct {
	number int
}

func (counter *Counter) Increment() {
	counter.number++
}

func (counter Counter) Value() int {
	return counter.number
}

func main() {
	counter := Counter{10} // 0
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.Value())

}
