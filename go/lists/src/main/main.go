package main

import (
	"fmt"
	"queue"
)

func main() {
	queue := queue.ArrayQueue{}

	queue.Init(10)
	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	for i := 0; i < 3; i++ {
		queue.Pop()
	}

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	peek, errorr := queue.Peek()
	fmt.Printf("Frente: %d", peek)
	fmt.Errorf("error : %s", errorr)

}
