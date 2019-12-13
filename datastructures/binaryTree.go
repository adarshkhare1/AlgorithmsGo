package datastructures

// package main

import (
	"strconv"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(root *TreeNode) *BinaryTree {
	return &BinaryTree{root: root}
}

// helper function for t.depth
func _calculateDepth(n *TreeNode, depth uint) uint {
	if n == nil {
		return depth
	}
	return _max(_calculateDepth(n.Left, depth+1), _calculateDepth(n.Right, depth+1))
}


func _max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func NewTreeNode(val int) *TreeNode {
	n := &TreeNode{val, nil, nil}
	return n
}

func Inorder(n *TreeNode, result string) string {
	if n == nil {
		return result
	}
	if n.Left != nil {
		result = Inorder(n.Left, result)
	}
	result += " " + strconv.Itoa(n.Value)
	if n.Right != nil {
		result = Inorder(n.Right, result)
	}
	return result
}


func (t *BinaryTree) Depth() uint {
	return _calculateDepth(t.root, 0)
}

