package linked

import "errors"

//LinkedQueue es la implementaci�n basada en nodos de la interfaz Queue.
type LinkedQueue struct {
	List *Node
}

//NewLinkedQueue crea una nueva instancia de LinkedQueue.
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{List: nil}
}

//IsEmpty regresa si la lista est� vac�a.
func (lq *LinkedQueue) IsEmpty() bool {
	return lq.List == nil
}

//Enqueue agrega un elemento a la cola.
func (lq *LinkedQueue) Enqueue(number int) {
	node := NewNode(number)
	if lq.IsEmpty() {
		lq.List = node
	} else if lq.List.Next == nil {
		lq.List.Next = node
	} else {
		current := lq.List
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

//Dequeue elimina el primer elemento de la lista y regresa su valor.
func (lq *LinkedQueue) Dequeue() (int, error) {
	if lq.IsEmpty() {
		return 0, errors.New("the list is empty")
	}

	value := lq.List.Data
	lq.List = lq.List.Next
	return value, nil
}

//First regresa el valor del primer elemento de la cola.
func (lq *LinkedQueue) Front() int {
	if lq.IsEmpty() {
		return 0
	}
	return lq.List.Data
}
