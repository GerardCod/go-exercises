package linked

import "fmt"

//Insert inserts a new node in a list.
func Insert(root *Node, data int) *Node {
	newNode := NewNode(data)
	if root == nil {
		root = newNode
		return root
	} else if root.Next == nil {
		root.Next = newNode
		return root
	} else {
		current := root
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	return root
}

//Show iterates over a 
func Show(root *Node) {
	if root == nil {
		fmt.Println("The list is empty")
	} else {
		current := root
		for current.Next != nil {
			fmt.Printf("%d ", current.Data)
			current = current.Next
		}
		fmt.Printf("%d", current.Data)
	}
}

//DeleteRepeated removes the repeated items in a linked list.
func DeleteRepeated(root *Node) *Node {
	if root == nil {
		return nil
	} else {
		current := root
		aux := root
		hash := make(map[int]bool)
		for current != nil {
			if !hash[current.Data] {
				hash[current.Data] = true
				aux = current
			} else {
				aux = deleteNode(aux, current.Next)
			}
			current = current.Next
		}
	}
	return root
}

func deleteNode(root *Node, newNextNode *Node) *Node {
	root.Next = newNextNode
	return root
}