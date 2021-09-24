package linked

import (
	"errors"
	"fmt"
)

//LinkedList es la implementacion por nodos de la interfaz lista.
type LinkedList struct {
  Head *Node
  Tail *Node
  Count int
}

//NewLinkedList crea una nueva instancia de LinkedList.
func NewLinkedList() *LinkedList {
  return &LinkedList{Head: nil, Tail: nil, Count: 0}
}

//IsEmpty regresa si la lista esta vacia.
func (ll *LinkedList) IsEmpty() bool {
  return ll.Count == 0
}

//Count regresa la cantidad de elementos que posee la lista.
func (ll *LinkedList) Size() int {
  return ll.Count
}

//First regresa el primer elemento de la lista.
func (ll *LinkedList) First() (int, error) {
  if ll.IsEmpty() {
    return 0, errors.New("the list is empty")
  }
  return ll.Head.Data, nil
}

//Last regresa el ultimo elemento de la lista.
func (ll *LinkedList) Last() (int, error) {
  if ll.IsEmpty() {
    return 0, errors.New("the list is empty")
  }

  if ll.Tail == nil {
    return 0, errors.New("the tail is nil")
  }

  return ll.Tail.Data, nil
}

//Append agrega un elemento al final de la lista.
func (ll *LinkedList) Append(number int) {
  node := NewNode(number)
  if ll.IsEmpty() {
    ll.Head = node
  } else if ll.Head.Next == nil {
    ll.Head.Next = node
    ll.Tail = ll.Head.Next
  } else {
    ll.Tail.Next = node
    ll.Tail = node
  }
  ll.Count++
}

//IndexOf regresa el indice de un elemento dentro de la lista enlazada. Regresa -1 en caso de que el numero no exista dentro de la lista.
func (ll *LinkedList) IndexOf(number int) int {
  index := -1
  if ll.Head.Data == number {
    index = 0
  } else if ll.Tail.Data == number {
    index = ll.Size() - 1
  } else {
    current := ll.Head
    for current != nil {
      index++
      if current.Data == number {
	      return index
      }
      current = current.Next
    }
    index = -1
  }
  return index
}

//Shift agrega un elemento al inicio de la lista.
func (ll *LinkedList) Shift(number int) {
  node := NewNode(number)
  if ll.IsEmpty() {
    ll.Head = node
  } else {
    node.Next = ll.Head
    ll.Head = node
  }
}

//Pop elimina el ultimo elemento de la lista y regresa su valor.
func (ll *LinkedList) Pop() (int, error) {
  if ll.IsEmpty() {
    return 0, errors.New("the list is empty")
  } else {
    current := ll.Head
    for current.Next != ll.Tail {
       current = current.Next
    }
    last := ll.Tail.Data
    ll.Tail = current
    ll.Count--
    return last, nil
  }
}

//Unshift elimina el primer elemento de la lista y regresa su valor
func (ll *LinkedList) Unshift() (int, error) {
  if ll.IsEmpty() {
    return 0, errors.New("the list is empty")
  } else {
    value := ll.Head.Data
    ll.Head = ll.Head.Next
    ll.Count--
    return value, nil
  }
}

//Show imprime todos los elementos de la lista separados por espacios.
func (ll *LinkedList) Show() error {
  if ll.IsEmpty() {
    return errors.New("the list is empty")
  } else {
     current := ll.Head
     for current != nil {
       fmt.Printf("%d ", current.Data)
       current = current.Next
     }
  }
  return nil
}
