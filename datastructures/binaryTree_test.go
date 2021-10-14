package datastructures

import (
	"Algorithms/testHelper"
	"testing"
)



func TestTreeTraversal(t *testing.T) {
	tree := getSampleTestTree()
	testHelper.VerifyUintsAreEqual(t, 6, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2 3 4 5 6 10", tree.Inorder())
	testHelper.VerifyStringsAreEqual(t, " 0 1 4 3 2 10 6 5", tree.Preorder())
	testHelper.VerifyStringsAreEqual(t, " 2 3 5 6 10 4 1 0", tree.Postorder())
}

func TestTreeInOrderDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyUintsAreEqual(t, 3, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2", tree.Inorder())
}

func TestTreePreOrderDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyUintsAreEqual(t, 3, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2", tree.Preorder())
}

func TestTreePostOrderDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyUintsAreEqual(t, 3, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 2 1 0", tree.Postorder())
}


func getSampleTestTree() *BinaryTree {
	tree := NewBinaryTree(0)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(2)

	return tree
}

func getSampleTreeWithDepth1() *BinaryTree {
	tree := NewBinaryTree(0)
	tree.Insert(1)
	tree.Insert(2)
	return tree
}
