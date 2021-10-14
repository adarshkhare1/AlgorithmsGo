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

// Inorder traversal print left node, root, right node order
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

// PostOrder traversal print left node, right node, root order
func PostOrder(n *TreeNode, result string) string {
	if n == nil {
		return result
	}
	if n.Left != nil {
		result = PostOrder(n.Left, result)
	}
	if n.Right != nil {
		result = PostOrder(n.Right, result)
	}
	result += " " + strconv.Itoa(n.Value)
	return result
}

// PreOrder traversal prints root, left node, right node order
func PreOrder(n *TreeNode, result string) string {
	if n == nil {
		return result
	}
	result += " " + strconv.Itoa(n.Value)
	if n.Left != nil {
		result = PreOrder(n.Left, result)
	}
	if n.Right != nil {
		result = PreOrder(n.Right, result)
	}
	return result
}


func (t *BinaryTree) Depth() uint {
	return _calculateDepth(t.root, 0)
}

