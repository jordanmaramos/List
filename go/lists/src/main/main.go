package main

import (
	"fmt"
	"list"
)

func main() {
	array := list.ArrayList{}
	array.Init()

	for i := 0; i < 10; i++ {
		array.Add(i)
	}
	fmt.Printf("Tamanho da lista: %d\n", array.Length())
	array.PrintList()

}
