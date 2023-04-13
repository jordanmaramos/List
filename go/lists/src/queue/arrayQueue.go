package queue

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	front    int
	back     int
	values   []int
	length   int
	capacity int
}

func (arrayQueue *ArrayQueue) Init(capacity int) error {
	if arrayQueue.length > 0 {
		return fmt.Errorf("arrayQueue already initialized")
	}

	if capacity > 0 {
		arrayQueue.values = make([]int, capacity)
		arrayQueue.front = -1
		arrayQueue.back = -1
		arrayQueue.length = 0
		arrayQueue.capacity = capacity
		return nil
	} else {
		return errors.New("invalid capacity")
	}
}

func (arrayQueue *ArrayQueue) Push(value int) {
	if arrayQueue.IsEmpty() {
		arrayQueue.front = 0
		arrayQueue.back = 0
		arrayQueue.values[0] = value
		arrayQueue.length = 1
	} else {
		arrayQueue.back = (arrayQueue.back + 1) % arrayQueue.capacity
		arrayQueue.values[arrayQueue.back] = value
		arrayQueue.length++
	}
}

func (arrayQueue *ArrayQueue) Pop() (int, error) {
	if arrayQueue.IsEmpty() {
		return -1, fmt.Errorf("empty queue")
	}

	aux := arrayQueue.values[arrayQueue.front]
	//If queue has only one element
	if arrayQueue.length == 1 {
		arrayQueue.front = -1
		arrayQueue.back = -1
		arrayQueue.length = 0
		return aux, nil
	}
	arrayQueue.front = (arrayQueue.front + 1) % arrayQueue.capacity
	arrayQueue.length--
	return aux, nil
}

func (arrayQueue *ArrayQueue) Peek() (int, error) {
	if arrayQueue.IsEmpty() {
		return -1, fmt.Errorf("empty queue")
	}
	return arrayQueue.values[arrayQueue.front], nil
}

func (arrayQueue *ArrayQueue) Size() int {
	return arrayQueue.length
}

func (arrayQueue *ArrayQueue) IsEmpty() bool {
	return arrayQueue.back == -1 && arrayQueue.front == -1 && arrayQueue.length == 0
}

func (arrayQueue *ArrayQueue) IsFull() bool {
	return ((arrayQueue.back + 1) % arrayQueue.capacity) == arrayQueue.front
}
