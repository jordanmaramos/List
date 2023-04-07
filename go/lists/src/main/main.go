package main

import (
	"list"
)

func main() {
	array := list.ArrayList{}
	array.Init()

	for i := 1; i < 10; i++ {
		array.Add(i)
	}
	array.PrintList()
}
