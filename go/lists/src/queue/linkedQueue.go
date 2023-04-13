package queue

import (
	"errors"
	"fmt"
)

type LinkedQueue struct {
	front  *Node
	back   *Node
	length int
}

type Node struct {
	value int
	next  *Node
}

func (linkedQueue *LinkedQueue) Init(i int) error {
	//TODO implement me
	return fmt.Errorf("")
}

func (linkedQueue *LinkedQueue) Push(value int) {
	newNode := &Node{value: value, next: nil}

	if linkedQueue.length == 0 {
		linkedQueue.front = newNode
		linkedQueue.back = newNode
	} else {
		linkedQueue.back.next = newNode
		linkedQueue.back = newNode
	}

	linkedQueue.length++
}

func (linkedQueue *LinkedQueue) Pop() (int, error) {
	if linkedQueue.length == 0 {
		return -1, errors.New("linkedQueue is empty")
	}

	value := linkedQueue.front.value
	linkedQueue.front = linkedQueue.front.next
	linkedQueue.length--

	return value, nil
}

func (linkedQueue *LinkedQueue) Peek() (int, error) {
	if linkedQueue.length == 0 {
		return -1, errors.New("linkedQueue is empty")
	}

	return linkedQueue.front.value, nil
}

func (linkedQueue *LinkedQueue) Size() int {
	return linkedQueue.length
}

func (linkedQueue *LinkedQueue) IsEmpty() bool {
	return linkedQueue.length == 0
}

func (linkedQueue *LinkedQueue) IsFull() bool {
	//TODO implement me
	return true
}
