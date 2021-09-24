package array

import "errors"

//ArrayQueue es la implementacion de arreglos de la interfaz cola.
type ArrayQueue struct {
	List  []int
	First int
}

//NewArrayQueue crea una nueva instancia de ArrayQueue
func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{List: []int{}, First: -1}
}

func (aq *ArrayQueue) IsEmpty() bool {
	return aq.First == -1 || aq.First == len(aq.List)
}

//Enqueue agrega un elemento a la cola.
func (aq *ArrayQueue) Enqueue(number int) {
	if aq.IsEmpty() {
		aq.First = 0
	}
	aq.List = append(aq.List, number)
}

//Dequeue elimina el primer elemento de la cola y devuelve su valor.
func (aq *ArrayQueue) Dequeue() (int, error) {
	if aq.IsEmpty() {
		return 0, errors.New("the list is empty")
	}
	value := aq.List[aq.First]
	aq.First++
	return value, nil
}

func (aq *ArrayQueue) Front() int {
	if aq.IsEmpty() {
		return 0
	}

	return aq.List[aq.First]
}
