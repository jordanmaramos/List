package main

import "fmt"

func BubbleSort(values []int) {
	size := len(values)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}

func InsertionSort(values []int) {
	size := len(values)
	for i := 0; i < size; i++ {
		for j := i; j > 0; j-- {
			if values[j] < values[j-1] {
				values[j-1], values[j] = values[j], values[j-1]
			} else {
				break
			}
		}
	}
}

func main() {
	//Example BubbleSort
	values := []int{64, 34, 25, 12, 22, 11, 8, 7, 6, 6, 4, 30}

	fmt.Println("Array original:", values)
	InsertionSort(values)
	fmt.Println("Array ordenado:", values)
}
