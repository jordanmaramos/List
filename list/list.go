package list

import "fmt"

type LinkedList struct {
	head      *Node
	lastIndex int
}

type Node struct {
	value int
	next  *Node
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func (linkedList *LinkedList) length() int {
	return linkedList.lastIndex + 1
}