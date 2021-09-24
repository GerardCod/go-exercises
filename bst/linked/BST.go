package linked

import "fmt"

//LinkedBST is the implementation with double linked list of BinarySearchTree interface.
type LinkedBST struct {
	Root   *Node
	Length int
}

//NewLinkedBST creates a new instance of LinkedBST.
func NewLinkedBST() *LinkedBST {
	return &LinkedBST{Root: nil}
}

//Add adds a new element in the BST.
func (lb *LinkedBST) Add(value int) {
	lb.Root = lb.addHelper(lb.Root, value)
}

func (lb *LinkedBST) addHelper(root *Node, value int) *Node {
	if root == nil {
		root = NewNode(value)
	} else {
		if root.Data < value {
			root.Right = lb.addHelper(root.Right, value)
		} else {
			root.Left = lb.addHelper(root.Left, value)
		}
	}
	return root
}

//InOrder iterates over the bst in order way.
func (lb *LinkedBST) InOrder() {
	lb.inOrderHelper(lb.Root)
}

func (lb *LinkedBST) inOrderHelper(node *Node) {
	if node == nil {
		return
	}
	lb.inOrderHelper(node.Left)
	fmt.Printf("%d ", node.Data)
	lb.inOrderHelper(node.Right)
}

//PreOrder goes through the BST in preorder way.
func (lb *LinkedBST) PreOrder() {
	lb.preOrderHelper(lb.Root)
}

func (lb *LinkedBST) preOrderHelper(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.Data)
	lb.preOrderHelper(node.Left)
	lb.preOrderHelper(node.Right)
}

//PostOrder goes through the BST in postorder way
func (lb *LinkedBST) PostOrder() {
	lb.postOrderHelper(lb.Root)
}

func (lb *LinkedBST) postOrderHelper(node *Node) {
	if node == nil {
		return
	}

	lb.postOrderHelper(node.Left)
	lb.postOrderHelper(node.Right)
	fmt.Printf("%d ", node.Data)
}

//LevelOrder goes through the BST in level order way
func (lb *LinkedBST) LevelOrder() {
	level := 1
	for lb.levelOrderHelper(lb.Root, level) {
		level++
	}
}

func (lb *LinkedBST) levelOrderHelper(node *Node, level int) bool {
	if node == nil {
		return false
	}

	if level == 1 {
		fmt.Printf("%d ", node.Data)
		return true
	}

	left := lb.levelOrderHelper(node.Left, level - 1)
	right := lb.levelOrderHelper(node.Right, level - 1)
	return left || right
}
