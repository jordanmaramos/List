package list

import (
	"fmt"
)

type ArrayList struct {
	values    []int
	lastIndex int
}

func (arrayList *ArrayList) Init() {
	arrayList.values = make([]int, 5)
	arrayList.lastIndex = 0
}

func (arrayList *ArrayList) Increment() {
	//Cria novo array vazio
	newArrayList := make([]int, len(arrayList.values)+1)
	//Preenche novo array com valores de arrayList
	for i := 0; i < len(arrayList.values); i++ {
		newArrayList[i] = arrayList.values[i]
	}
	//Devolve novo array para arrayList
	arrayList.values = newArrayList
}

func (arrayList *ArrayList) Add(value int) {
	if arrayList.lastIndex == len(arrayList.values) {
		arrayList.Increment()
	}
	arrayList.values[arrayList.lastIndex] = value
	arrayList.lastIndex++
}

func (arrayList *ArrayList) AddOnIndex(value int, index int) {
	if arrayList.lastIndex == len(arrayList.values) {
		arrayList.Increment()
	}
	for i := arrayList.lastIndex; i > index; i-- {
		arrayList.values[i] = arrayList.values[i-1]
	}
	arrayList.values[index] = value
	arrayList.lastIndex++
}

func (arrayList *ArrayList) Remove() error {
	if arrayList.lastIndex > 0 {
		arrayList.lastIndex--
		return nil
	}
	return fmt.Errorf("Array vazio. Não há o que remover")
}

func (arrayList *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < arrayList.lastIndex {
		for i := index; i < arrayList.lastIndex; i++ {
			arrayList.values[i] = arrayList.values[i+1]
		}
		arrayList.lastIndex--
		return nil
	}

	return fmt.Errorf("Array vazio. Não há o que remover")
}

func (arrayList *ArrayList) Get(index int) int {
	if index >= 0 && index < arrayList.lastIndex {
		return arrayList.values[index]
	}
	return -1
}

func (arrayList *ArrayList) Set(value int, index int) {
	if index >= 0 && index < arrayList.lastIndex {
		arrayList.values[index] = value
	}
}

func (arrayList *ArrayList) Length() int {
	return arrayList.lastIndex
}

func (arrayList *ArrayList) PrintList() {
	length := arrayList.Length()
	for i := 0; i < length; i++ {
		fmt.Print(arrayList.values[i])
		fmt.Print(" ")
	}
}
