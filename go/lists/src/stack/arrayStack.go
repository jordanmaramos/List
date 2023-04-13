package stack

import (
	"errors"
)

type ArrayStack struct {
	values []int
	length int
}

func (arrayStack *ArrayStack) Init(capacity int) error {
	if arrayStack.length > 0 {
		return errors.New("stack already initialized")
	}

	if capacity > 0 {
		arrayStack.values = make([]int, capacity)
		arrayStack.length = 0
	} else {
		return errors.New("invalid capacity value")
	}

	return nil
}

func (arrayStack *ArrayStack) Double() {
	newArrayStack := make([]int, len(arrayStack.values)*2)
	for i := 0; i < len(arrayStack.values); i++ {
		newArrayStack[i] = arrayStack.values[i]
	}
	arrayStack.values = newArrayStack
}

func (arrayStack *ArrayStack) Push(value int) {
	if arrayStack.length == len(arrayStack.values) {
		arrayStack.Double()
	}
	arrayStack.values[arrayStack.length] = value
	arrayStack.length++
}

func (arrayStack *ArrayStack) Pop() (int, error) {
	if arrayStack.length > 0 {
		item := arrayStack.values[arrayStack.length-1]
		arrayStack.length--

		return item, nil
	} else {
		return -1, errors.New("arrayStack is empity")
	}
}

func (arrayStack *ArrayStack) Peek() (int, error) {
	if arrayStack.length > 0 {
		return arrayStack.values[arrayStack.length-1], nil
	} else {
		return -1, errors.New("arrayStack is empty")
	}
}

func (arrayStack *ArrayStack) Empty() bool {
	return arrayStack.length == 0
}

func (arrayStack *ArrayStack) Size() int {
	return arrayStack.length
}
