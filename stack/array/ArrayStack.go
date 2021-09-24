package array

import "errors"

//ArrayStack is the implementation with arrays of Stack interface.
type ArrayStack struct {
	List  []int
	First int
}

//NewArrayStack creates a new instance of ArrayStack.
func NewArrayStack() *ArrayStack {
	return &ArrayStack{List: []int{}, First: -1}
}

//Push adds a new element in the stack
func (as *ArrayStack) Push(number int) {
	as.First++
	as.List = append(as.List, number)
}

//Pop elimina el primer elemento de la lista y regresa el valor del elemento eliminado.
func (as *ArrayStack) Pop() (int, error) {
	if as.IsEmpty() {
		return 0, errors.New("the list is empty")
	}

	value := as.List[as.First]
	as.First--
	return value, nil
}

//IsEmpty regresa si la lista esta vacia.
func (as *ArrayStack) IsEmpty() bool {
	return as.First == -1
}

//First regresa el primer elemento de la pila
func (as *ArrayStack) Peek() (int, error) {
	if as.IsEmpty() {
		return 0, errors.New("the list is empty")
	}
	return as.List[as.First], nil
}
