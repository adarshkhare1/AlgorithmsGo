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

// helper function for t.depth
func _calculate_depth(n *TreeNode, depth int) int {
	if n == nil {
		return depth
	}
	return _max(_calculate_depth(n.Left, depth+1), _calculate_depth(n.Right, depth+1))
}


func _max(a, b int) int {
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


func (t *BinaryTree) Depth() int {
	return _calculate_depth(t.root, 0)
}

