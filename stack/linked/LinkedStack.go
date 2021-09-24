package linked

import "errors"

//LinkedStack is the implementation of Stack interface.
type LinkedStack struct {
	Head *Node
}

//NewLinkedStack creates a new instance of LinkedStack.
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{Head: nil}
}

//Push adds a new item at the end of the stack.
func (ls *LinkedStack) Push(number int) {
	node := NewNode(number)
	if ls.Head == nil {
		ls.Head = node
	} else {
		node.Next = ls.Head
		ls.Head = node
	}
}

//Pop removes the first element of the stack and retrieves its value.
func (ls *LinkedStack) Pop() (int, error) {
	if ls.IsEmpty() {
		return 0, errors.New("the list is empty")
	}

	value := ls.Head.Data
	ls.Head = ls.Head.Next
	return value, nil
}

//IsEmpty returns if the stack is empty.
func (ls *LinkedStack) IsEmpty() bool {
	return ls.Head == nil
}

//Peek returns the value of the first element of the stack.
func (ls *LinkedStack) Peek() (int, error) {
	if ls.IsEmpty() {
		return 0, errors.New("the list is empty")
	}
	return ls.Head.Data, nil
}
