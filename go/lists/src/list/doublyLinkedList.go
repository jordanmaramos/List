package list

type DoublyLinkedList struct {
	head      *NodeDoubly
	tail      *NodeDoubly
	lastIndex int
}

type NodeDoubly struct {
	previous *NodeDoubly
	next     *NodeDoubly
	value    int
}

func (doublyLinkedList *DoublyLinkedList) Init() {
	doublyLinkedList.head = nil
	doublyLinkedList.tail = nil
	doublyLinkedList.lastIndex = -1
}

func (doublyLinkedList *DoublyLinkedList) Add(value int) {
	newNode := NodeDoubly{value: value, next: nil, previous: nil}

	lastNode := doublyLinkedList.head //Nó atual inicia pela cabeça
	previousNode := lastNode          //Nó anterior do próximo nó, é o nó atual

	//Enquanto o último nó tiver next != nil
	for lastNode != nil {
		previousNode = lastNode  //Nó anterior do próximo nó aponta para nó atual
		lastNode = lastNode.next //Nó atual recebe próximo nó do nó atual
	}

	//Lista vazia, primeira iteração
	if previousNode == nil {
		doublyLinkedList.head = &newNode
		doublyLinkedList.tail = &newNode
	} else {
		/* Encontrei último nó (previousNode), cujo o next é nil,
		faço seu next apontar para o novo nó, e o previous do novo
		nó apontar para o antigo último nó (previousNode)
		*/
		previousNode.next = &newNode
		newNode.previous = previousNode
	}

	//Incremento última posição
	doublyLinkedList.lastIndex++
}

func (doublyLinkedList *DoublyLinkedList) AddOnIndex(value, index int) {
	newNode := NodeDoubly{value: value, next: nil, previous: nil}

	lastNode := doublyLinkedList.head //Nó atual inicia pela cabeça
	previousNode := lastNode          //Nó anterior do próximo nó, é o nó atual

	//Enquanto o último nó tiver next != nil
	for lastNode != nil {
		previousNode = lastNode  //Nó anterior do próximo nó aponta para nó atual
		lastNode = lastNode.next //Nó atual recebe próximo nó do nó atual
	}

	//Lista vazia, primeira iteração
	if previousNode == nil {
		doublyLinkedList.head = &newNode
		doublyLinkedList.tail = &newNode
	} else {
		/* Encontrei último nó (previousNode), cujo o next é nil,
		faço seu next apontar para o novo nó, e o previous do novo
		nó apontar para o antigo último nó (previousNode)
		*/
		previousNode.next = &newNode
		newNode.previous = previousNode
	}

	//Incremento última posição
	doublyLinkedList.lastIndex++
}
