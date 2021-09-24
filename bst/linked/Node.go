package linked

//Node is a double linked list for the BST.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

//NewNode creates a new instance of Node struct.
func NewNode(value int) *Node {
	return &Node{Data: value, Left: nil, Right: nil}
}
