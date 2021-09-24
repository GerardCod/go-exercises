package linked

//Node representa un nodo de una lista enlazada o cualquier estructura de datos basada en nodos.
type Node struct {
	Data int
	Next *Node
}

//NewNode crea una nueva instancia de Node.
func NewNode(number int) *Node {
	return &Node{Data: number, Next: nil}
}
