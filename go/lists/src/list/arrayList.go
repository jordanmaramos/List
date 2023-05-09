package list

import (
	"fmt"
)

type ArrayList struct {
	Values    []int
	LastIndex int
}

func (arrayList *ArrayList) Init() {
	arrayList.Values = make([]int, 5)
	arrayList.LastIndex = 0
}

func (arrayList *ArrayList) Increment() {
	//Cria novo array vazio
	newArrayList := make([]int, len(arrayList.Values)+1)
	//Preenche novo array com valores de arrayList
	for i := 0; i < len(arrayList.Values); i++ {
		newArrayList[i] = arrayList.Values[i]
	}
	//Devolve novo array para arrayList
	arrayList.Values = newArrayList
}

func (arrayList *ArrayList) Add(value int) {
	if arrayList.LastIndex == len(arrayList.Values) {
		arrayList.Increment()
	}
	arrayList.Values[arrayList.LastIndex] = value
	arrayList.LastIndex++
}

func (arrayList *ArrayList) AddOnIndex(value int, index int) {
	if arrayList.LastIndex == len(arrayList.Values) {
		arrayList.Increment()
	}
	for i := arrayList.LastIndex; i > index; i-- {
		arrayList.Values[i] = arrayList.Values[i-1]
	}
	arrayList.Values[index] = value
	arrayList.LastIndex++
}

func (arrayList *ArrayList) Remove() error {
	if arrayList.LastIndex > 0 {
		arrayList.LastIndex--
		return nil
	}
	return fmt.Errorf("Array vazio. Não há o que remover")
}

func (arrayList *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < arrayList.LastIndex {
		for i := index; i < arrayList.LastIndex; i++ {
			arrayList.Values[i] = arrayList.Values[i+1]
		}
		arrayList.LastIndex--
		return nil
	}

	return fmt.Errorf("Array vazio. Não há o que remover")
}

func (arrayList *ArrayList) Get(index int) int {
	if index >= 0 && index < arrayList.LastIndex {
		return arrayList.Values[index]
	}
	return -1
}

func (arrayList *ArrayList) Set(value int, index int) {
	if index >= 0 && index < arrayList.LastIndex {
		arrayList.Values[index] = value
	}
}

func (arrayList *ArrayList) Length() int {
	return arrayList.LastIndex
}

func (arrayList *ArrayList) PrintList() {
	length := arrayList.Length()
	for i := 0; i < length; i++ {
		fmt.Print(arrayList.Values[i])
		fmt.Print(" ")
	}
}
