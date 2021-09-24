package linked

//Node is the node
type Node struct {
	Data int
	Next *Node
}

//NewNode creates a new instance of Node.
func NewNode(number int) *Node {
	return &Node{Data: number, Next: nil}
}