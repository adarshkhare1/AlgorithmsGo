package redBlackTree

import (
	"AlgorithmsGo/testHelper"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestRBTreeTraversal(t *testing.T) {
	tree := getSampleTestTree()
	testHelper.VerifyIntsAreEqual(t, 10, tree.Root.TreeMaximum().Key)
	testHelper.VerifyIntsAreEqual(t, 0, tree.Root.TreeMinimum().Key)
	testHelper.VerifyUintsAreEqual(t, 4, tree.Depth())
	testHelper.VerifyStringsAreEqual(t, " 0 1 2 3 4 5 6 10", tree.Inorder())
}

func TestTreeTraversalDepth1(t *testing.T) {
	tree := getSampleTreeWithDepth1()
	testHelper.VerifyIntsAreEqual(t, 2, tree.Root.TreeMaximum().Key)
	testHelper.VerifyIntsAreEqual(t, 0, tree.Root.TreeMinimum().Key)
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

func TestRBTreeLargeSizeWithRandomNumbers(t *testing.T) {
	nCount := 512
	maxKeyValue := 100000
	tree := buildATreeWithRandomNumbers(nCount, maxKeyValue)
	nums := getNodeValuesArray(tree.Inorder())
	testHelper.VerifyIntsAreEqual(t, nCount, len(nums))
	testHelper.VerifyArrayIsSortedAscending(t, nums)
	fmt.Printf("Tree Depth is %d for %d nodes.\n", tree.Depth(), nCount)
}

func TestRBTreeLargeSizeRandomDeletion(t *testing.T) {
	nCount := 1024
	maxKeyValue := 100000
	tree := buildATreeWithRandomNumbers(nCount, maxKeyValue)
	nums := getNodeValuesArray(tree.Inorder())
	//Delete some nodes randomly
	countToDelete := 64
	deleteTreeNodesRandomly(tree, nums, countToDelete)
	nums = getNodeValuesArray(tree.Inorder())
	testHelper.VerifyIntsAreEqual(t, nCount-countToDelete+1, len(nums))
	testHelper.VerifyArrayIsSortedAscending(t, nums)
}

func buildATreeWithRandomNumbers(nodeCount int, maxKeyValue int) *RedBlackTree {
	tree := NewRedBlackBinaryTreeEmpty()
	for i := 0; i < nodeCount; i++ {
		n := rand.Intn(maxKeyValue)
		tree.Insert(n)
	}
	return tree
}

func deleteTreeNodesRandomly(tree *RedBlackTree, nums []int, countToDelete int) {
	nCount := len(nums)
	if countToDelete > (nCount - countToDelete) {
		panic("Cannot delete more nodes than we have.")
	}
	for i := 0; i < countToDelete; i++ {
		n := rand.Intn(nCount - countToDelete) // just to be safe
		nodeToDelete := SearchTreeNode(tree.Root, nums[n])
		tree.Delete(nodeToDelete)
	}
}

func getNodeValuesArray(str string) []int {
	strs := strings.Split(str, " ")
	nums := make([]int, len(strs)-1)
	for i := range nums {
		if strings.TrimSpace(strs[i]) != "" {
			nums[i], _ = strconv.Atoi(strs[i])
		}
	}
	return nums
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
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(0)
	return tree
}
