package redBlackTree

import "math"

type Color int8

const (
	Red   Color = 0
	Black       = 1
)

type RedBlackTreeNode struct {
	Key  int
	Left *RedBlackTreeNode
	Right  *RedBlackTreeNode
	Parent *RedBlackTreeNode
	Color  Color
	IsNil  bool
}

func _newRedBlackTreeNode(k int) *RedBlackTreeNode {
	n := &RedBlackTreeNode{
		Key:    k,
		Parent: _nilRedBlackTreeNode(nil),
		Color:  Red,
		IsNil:  false,
	}
	n.Left = _nilRedBlackTreeNode(n)
	n.Right = _nilRedBlackTreeNode(n)
	return n
}

func _nilRedBlackTreeNode(parent *RedBlackTreeNode) *RedBlackTreeNode {
		return &RedBlackTreeNode{
			Key:    math.MinInt64,
			Left:   nil,
			Right:  nil,
			Parent: parent,
			Color:  Black,
			IsNil:  true,
		}
}

func (n *RedBlackTreeNode)_setLeft(node *RedBlackTreeNode){
	if n == node {
		panic("Children cannot be set to itself.")
	}
	if !n.IsNil{
		n.Left = node
		node.Parent = n
	}else {
		panic("Children for nilNode is not allowed.")
	}
}

func (n *RedBlackTreeNode)_setRight(node *RedBlackTreeNode){
	if n == node {
		panic("Children cannot be set to itself.")
	}
	if !n.IsNil{
		n.Right = node
		node.Parent = n
	}else {
		panic("Children for nilNode is not allowed.")
	}

}

