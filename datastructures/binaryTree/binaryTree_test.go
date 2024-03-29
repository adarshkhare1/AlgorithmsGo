package binaryTree

import (
	"AlgorithmsGo/testHelper"
	"testing"
)

func TestTreeTraversal(t *testing.T) {
	tree := getSampleTestTree()
	testHelper.VerifyIntsAreEqual(t, treeMaximum(tree.Root).Key, 10)
	testHelper.VerifyIntsAreEqual(t, treeMinimum(tree.Root).Key, 0)
	testHelper.VerifyUintsAreEqual(t, 6, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2 3 4 5 6 10", tree.Inorder())
	testHelper.VerifyStringsAreEqual(t, " 0 1 4 3 2 10 6 5", tree.Preorder())
	testHelper.VerifyStringsAreEqual(t, " 2 3 5 6 10 4 1 0", tree.Postorder())
}

func TestTreeTraversalDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyIntsAreEqual(t, treeMaximum(tree.Root).Key, 2)
	testHelper.VerifyIntsAreEqual(t, treeMinimum(tree.Root).Key, 0)
	testHelper.VerifyUintsAreEqual(t, 2, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2", tree.Inorder())
	testHelper.VerifyStringsAreEqual(t, " 1 0 2", tree.Preorder())
	testHelper.VerifyStringsAreEqual(t, " 0 2 1", tree.Postorder())
}

func TestEmptyTree(t *testing.T) {
	tree := NewBinaryTreeEmpty()
	testHelper.VerifyStringsAreEqual(t, "", tree.Inorder())
	testHelper.VerifyStringsAreEqual(t, "", tree.Preorder())
	testHelper.VerifyStringsAreEqual(t, "", tree.Postorder())
	tree.Insert(0)
	testHelper.VerifyStringsAreEqual(t, " 0", tree.Inorder())
	testHelper.VerifyStringsAreEqual(t, " 0", tree.Preorder())
	testHelper.VerifyStringsAreEqual(t, " 0", tree.Postorder())
}

func TestTreeDeleteRoot(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	tree.Delete(tree.Root)
	testHelper.VerifyStringsAreEqual(t, " 0 2", tree.Inorder())
}

func TestTreeDeleteLeftOfRoot(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	tree.Delete(tree.Root.Left)
	testHelper.VerifyStringsAreEqual(t, " 1 2", tree.Inorder())
}

func TestTreeDeleteRightOfRoot(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	tree.Delete(tree.Root.Right)
	testHelper.VerifyStringsAreEqual(t, " 0 1", tree.Inorder())
}

func TestTreeDeleteNodes(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	tree.Delete(SearchTreeNode(tree.Root, 2))
	testHelper.VerifyStringsAreEqual(t, " 0 1", tree.Inorder())
	tree.Delete(SearchTreeNode(tree.Root, 0))
	testHelper.VerifyStringsAreEqual(t, " 1", tree.Inorder())
	tree.Delete(SearchTreeNode(tree.Root, 1))
	testHelper.VerifyStringsAreEqual(t, "", tree.Inorder())
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
	tree := NewBinaryTreeEmpty()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(0)
	return tree
}
