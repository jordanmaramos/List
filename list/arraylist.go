package list

import(
	"fmt"
)

type ArrayList struct {
	values []int
	lastIndex int
}

func (arrayList *ArrayList) init(){
	arrayList.values = make([]int, 5)
}

func (arrayList *ArrayList) increment(){
	//Cria novo array vazio
	newArrayList := make([]int, len(arrayList.values)+1)
	//Preenche novo array com valores de arrayList
	for i:=0; i < len(arrayList.values); i++{
		newArrayList[i] = arrayList.values[i]
	}
	//Devolve novo array para arrayList
	arrayList.values = newArrayList
}

func (arrayList *ArrayList) add(value int){
	if arrayList.lastIndex == len(arrayList.values){
		arrayList.increment()
	}
	arrayList.values[arrayList.lastIndex] = value
	arrayList.lastIndex++
}

func (arrayList *ArrayList) addOnIndex(value int, index int){
	if arrayList.lastIndex == len(arrayList.values){
		arrayList.increment()
	}
	for i := arrayList.lastIndex; i > index; i--{
		arrayList.values[i] = arrayList.values[i-1]
	} 
}

func (arrayList *ArrayList) remove() error{
	if arrayList.lastIndex > 0{
		arrayList.lastIndex--
	} 
	return fmt.Errorf("Array vazio. Não há o que remover")
}

func (arrayList *ArrayList) removeOnIndex(index int){
	if index >= 0 && index < arrayList.lastIndex{
		for i := index; i < arrayList.lastIndex; i++{
			arrayList.values[i] = arrayList.values[i+1]
		} 
		arrayList.lastIndex--
	}
}

func (arrayList *ArrayList) get(index int) int{
	if index >= 0 && index < arrayList.lastIndex{
		return arrayList.values[index]
	}
	return -1
}

func (arrayList *ArrayList) set(value int, index int){
	if index >= 0 && index < arrayList.lastIndex{
		arrayList.values[index] = value
	}
}

func (arrayList *ArrayList) length() int{
	return arrayList.lastIndex
}