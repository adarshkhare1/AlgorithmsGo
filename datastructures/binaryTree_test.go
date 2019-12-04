package datastructures

import (
	"testing"
)

func verifyBinaryTree(t *testing.T, tree BinaryTree, expected string){
	/* read the list until head is not nil */
	verifyString(t, expected, Inorder(tree.root, ""))
}

func TestNewTreeNode(t *testing.T) {
	tree := BinaryTree{nil}
	tree.root = NewTreeNode(0)
	tree.root.Left = NewTreeNode(1)
	tree.root.Right = NewTreeNode(2)
	tree.root.Left.Left = NewTreeNode(3)
	tree.root.Left.Right = NewTreeNode(4)
	tree.root.Right.Left = NewTreeNode(5)
	tree.root.Right.Right = NewTreeNode(6)
	tree.root.Right.Right.Right = NewTreeNode(10)
	verifyBinaryTree(t, tree, " 3 1 4 0 5 2 6 10")
	verifyInt(t, 4,  tree.Depth())
}

