package linked

//Node is a struct representing a node of a linked list
type Node struct {
	Data int
	Next *Node
}

//NewNode creates a new 
func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}
