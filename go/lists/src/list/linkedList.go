package list

import (
	"fmt"
)

type LinkedList struct {
	head      *Node
	lastIndex int
}

type Node struct {
	value int
	next  *Node
}

func (linkedList *LinkedList) Init() {
	linkedList.head = nil
	linkedList.lastIndex = -1
}

func (linkedList *LinkedList) Add(value int) {
	newNode := Node{value: value, next: nil}

	//Se é o primeiro nó da lista/lista vazia
	if linkedList.head == nil {
		linkedList.head = &newNode
	} else {
		//Caso não seja primeiro nó da lista, busca último nó
		lastNode := linkedList.head

		//Enquanto não encontra último nó, continua procurando
		for lastNode.next != nil {
			lastNode = lastNode.next
		}

		//Faço último nó apontar para o novo nó
		lastNode.next = &newNode

		//Incrementa posição do último nó
		linkedList.lastIndex++
	}
}

func (linkedList *LinkedList) AddOnIndex(value int, index int) error {
	//Verifica se o index é válido
	if index < 0 || index > linkedList.lastIndex {
		return fmt.Errorf("Index inválido")
	}

	newNode := &Node{value: value, next: nil}

	//Caso index == 0, apenas altero o seu valor
	if index == 0 {
		newNode.next = linkedList.head
		linkedList.head = newNode
	} else {
		//Caso index > 0, busco pelo nó
		previousNode := linkedList.head
		for i := 0; i < index-1; i++ {
			previousNode = previousNode.next
		}

		//Encontrei o nó, sai do laço
		newNode.next = previousNode.next
		previousNode.next = newNode
	}

	//Atualiza index do último nó
	linkedList.lastIndex++
	return nil
}

func (linkedList *LinkedList) Remove() error {
	//Se a lista estiver vazia, não há o que remover
	if linkedList.head == nil {
		return fmt.Errorf("Lista vazia")
	}

	//Se a lista só tem um elemento, apenas o removo
	if linkedList.head.next == nil {
		linkedList.head = nil
		linkedList.lastIndex--
	} else {
		//Se a lista possui mais de um elemento, busco o último para remoção
		previousNode := linkedList.head
		for i := 0; i < linkedList.lastIndex-1; i++ {
			previousNode = previousNode.next
		}

		//Cheguei ao penúltimo nó, removo o último nó
		previousNode.next = nil
	}

	//Atualiza index do último nó
	linkedList.lastIndex--
	return nil
}

func (linkedList *LinkedList) RemoveOnIndex(index int) error {
	//Se a lista estiver vazia, não há o que remover
	if linkedList.head == nil {
		return fmt.Errorf("Lista vazia")
	}

	if index <= 0 || index > linkedList.lastIndex {
		return fmt.Errorf("Index inválido")
	}

	//Se for o primeiro nó a ser removido
	if index == 0 {
		linkedList.head = linkedList.head.next
	} else {
		previousNode := linkedList.head
		for i := 0; i < index-1; i++ {
			previousNode = previousNode.next
		}

		//Se o index é na última posição, apenas removo
		if index == linkedList.lastIndex {
			previousNode.next = nil
		} else {
			//Caso não seja na última posição, aponto o next para o próximo node
			previousNode.next = previousNode.next.next
		}
	}

	//Atualiza index do último nó
	linkedList.lastIndex--
	return nil
}

func (linkedList *LinkedList) GetNode(index int) (*Node, error) {
	if linkedList.head == nil {
		return nil, fmt.Errorf("Lista vazia.")
	}

	if index < 0 || index > linkedList.lastIndex {
		return nil, fmt.Errorf("Index inválido.")
	}

	if index == 0 {
		return linkedList.head, nil
	} else {
		//Busca por node
		node := linkedList.head
		for i := 0; i < index; i++ {
			node = node.next
		}

		//Encontrei node, retorno
		return node, nil
	}

	return nil, nil
}

func (linkedList *LinkedList) GetValue(index int) (int, error) {
	if linkedList.head == nil {
		return 0, fmt.Errorf("Lista vazia.")
	}

	if index < 0 || index > linkedList.lastIndex {
		return 0, fmt.Errorf("Index inválido.")
	}

	node, erro := linkedList.GetNode(index)
	return node.value, erro
}

func (linkedList *LinkedList) SetValue(value int, index int) error {
	if linkedList.head == nil {
		return fmt.Errorf("Lista vazia.")
	}

	if index < 0 || index > linkedList.lastIndex {
		return fmt.Errorf("Index inválido.")
	}

	if index == 0 {
		linkedList.head.value = value
	} else {
		node := linkedList.head
		for i := 0; i < index; i++ {
			node = node.next
		}
		node.value = value
	}

	// node, erro := linkedList.getNode(index)
	// node.value = value
	return nil
}

func (linkedList *LinkedList) Length() int {
	return linkedList.lastIndex + 1
}
