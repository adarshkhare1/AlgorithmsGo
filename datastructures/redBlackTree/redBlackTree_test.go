package redBlackTree

import (
	"Algorithms/testHelper"
	"testing"
)

func TestRBTreeTraversal(t *testing.T) {
	tree := getSampleTestTree()
	testHelper.VerifyIntsAreEqual(t, 10, tree.Root.TreeMaximum().Value)
	testHelper.VerifyIntsAreEqual(t, 0, tree.Root.TreeMinimum().Value)
	testHelper.VerifyUintsAreEqual(t, 4, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2 3 4 5 6 10", tree.Inorder())
}

func TestTreeTraversalDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyIntsAreEqual(t, 2, tree.Root.TreeMaximum().Value)
	testHelper.VerifyIntsAreEqual(t, 0, tree.Root.TreeMinimum().Value)
	testHelper.VerifyUintsAreEqual(t, 2, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2", tree.Inorder())
}

func TestEmptyRBTree(t *testing.T) {
	tree := NewRedBlackBinaryTreeEmpty()
	testHelper.VerifyStringsAreEqual(t, "", tree.Inorder())
	tree.Insert(0)
	testHelper.VerifyStringsAreEqual(t, " 0", tree.Inorder())
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

func getSampleTestTree() *RedBlackTree {
	tree := NewRedBlackTree(0)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(3)

	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(10)
	return tree
}

func getSampleTreeWithDepth1() *RedBlackTree {
	tree := NewRedBlackBinaryTreeEmpty()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(0)
	return tree
}
