package datastructures

import (
	"Algorithms/testHelper"
	"testing"
)



func TestTreeInOrder(t *testing.T) {
	tree := getSampleTestTree()
	testHelper.VerifyUintsAreEqual(t, 4, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 3 1 4 0 5 2 6 10", Inorder(tree.root, ""))
}

func TestTreePreOrder(t *testing.T) {
	tree := getSampleTestTree()
	testHelper.VerifyUintsAreEqual(t, 4, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 3 4 2 5 6 10", PreOrder(tree.root, ""))
}

func TestTreePostOrder(t *testing.T) {
	tree := getSampleTestTree()
	testHelper.VerifyUintsAreEqual(t, 4, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 3 4 1 5 10 6 2 0", PostOrder(tree.root, ""))
}

func TestTreeInOrderDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyUintsAreEqual(t, 2, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 1 0 2", Inorder(tree.root, ""))
}

func TestTreePreOrderDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyUintsAreEqual(t, 2, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2", PreOrder(tree.root, ""))
}

func TestTreePostOrderDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyUintsAreEqual(t, 2, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 1 2 0", PostOrder(tree.root, ""))
}


func getSampleTestTree() *BinaryTree {
	tree := NewBinaryTree(NewTreeNode(0))
	tree.root.Left = NewTreeNode(1)
	tree.root.Right = NewTreeNode(2)
	tree.root.Left.Left = NewTreeNode(3)
	tree.root.Left.Right = NewTreeNode(4)
	tree.root.Right.Left = NewTreeNode(5)
	tree.root.Right.Right = NewTreeNode(6)
	tree.root.Right.Right.Right = NewTreeNode(10)
	return tree
}

func getSampleTreeWithDepth1() *BinaryTree {
	tree := NewBinaryTree(NewTreeNode(0))
	tree.root.Left = NewTreeNode(1)
	tree.root.Right = NewTreeNode(2)
	return tree
}

func verifyBinaryTree(t *testing.T, tree *BinaryTree, expected string){
	/* read the list until head is not nil */
	testHelper.VerifyStringsAreEqual(t, expected, Inorder(tree.root, ""))
}

