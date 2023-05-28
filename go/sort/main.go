package main

import "fmt"

func BubbleSort(values []int) {
	size := len(values)
	for i := 0; i < size; i++ {
		swapped := false
		for j := 0; j < size-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				swapped = true
			}
		}
		if !swapped {
			return
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

func SelectionSort(values []int) {
	n := len(values)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if values[j] < values[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			values[i], values[minIndex] = values[minIndex], values[i]
		}
	}
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func QuickSort(values []int, ini, fim int) {
	if fim > ini {
		pivot := partition(values, ini, fim)
		QuickSort(values, ini, pivot-1)
		QuickSort(values, pivot+1, fim)
	}
}

func partition(values []int, ini, fim int) int {
	pivot := values[fim]
	pivotIndex := 0
	for i := 0; i < fim; i++ {
		if values[i] <= pivot {
			values[i], values[pivotIndex] = values[pivotIndex], values[i]
			pivotIndex++
		}
	}
	values[fim], values[pivotIndex] = values[pivotIndex], values[fim]
	return pivotIndex
}

func CountingSort(values []int) []int {
	// 1 - Descobrir maior e menor valor
	menor := values[0]
	maior := values[0]
	for i := 0; i < len(values); i++ {
		if values[i] < menor {
			menor = values[i]
		} else if values[i] > maior {
			maior = values[i]
		}
	}

	// 2 - Vetor de contagem e 3 - contar ocorrências de cada elemento
	sizeVet := maior - menor + 1
	vetCont := make([]int, sizeVet)
	for i := 0; i < len(values); i++ {
		vetCont[values[i]-menor]++
	}
	// 4 - Soma cumulativa
	for i := 1; i < sizeVet; i++ {
		vetCont[i] += vetCont[i-1]
	}

	// 5 - Criar novo vetor
	vetOrd := make([]int, len(values))

	// 6 - Mapear elementos de values para vetOrd utilizando vetCont
	for i := 0; i < len(values); i++ {
		vetOrd[vetCont[values[i]-menor]-1] = values[i]
		vetCont[values[i]-menor]-- //Decremento a contagem daquele vaor, indicando que diminuiu a quantidade de vezes que aparece
	}
	return vetOrd
}

func main() {

	values := []int{10, 8, 6, 7, 3, 2, 1, 4, 5, 11, 12, 4, 4, 4}

	fmt.Println("Array original:", values)
	values = CountingSort(values)
	fmt.Println("Array ordenado:", values)
}
