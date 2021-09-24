package bst

//BinarySearchTree interface defines the methods of a BST.
type BinarySearchTree interface {
	Add(value int)
	InOrder()
	PreOrder()
	PostOrder()
	LevelOrder()
}
