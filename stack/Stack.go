package stack

//Stack defines the methods to create a stack data structure.
type Stack interface {
	Push(number int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
}
