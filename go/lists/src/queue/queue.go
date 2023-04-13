package queue

type IQueue interface {
	Init(int) error
	Push(int)          //Enqueue
	Pop() (int, error) //Dequeue
	Peek() (int, error)
	Size() int
	IsEmpty() bool
	IsFull() bool
}
