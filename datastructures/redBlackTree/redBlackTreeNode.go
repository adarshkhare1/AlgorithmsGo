package redBlackTree

import "math"

type Color int8

const (
	Red   Color = 0
	Black       = 1
)

type RedBlackTreeNode struct {
	Value  int
	Left   *RedBlackTreeNode
	Right  *RedBlackTreeNode
	Parent *RedBlackTreeNode
	Color  Color
	IsNil  bool
}

func _newRedBlackTreeNode(value int) *RedBlackTreeNode {
	n := &RedBlackTreeNode{
		Value:  value,
		Parent: nil,
		Color:  Red,
		IsNil:  false,
	}
	n.Left = _nilRedBlackTreeNode(n)
	n.Right = _nilRedBlackTreeNode(n)
	return n
}

func _nilRedBlackTreeNode(parent *RedBlackTreeNode) *RedBlackTreeNode {
		return &RedBlackTreeNode{
			Value:  math.MinInt64,
			Left:   nil,
			Right:  nil,
			Parent: parent,
			Color:  Black,
			IsNil:  true,
		}
}

func (n *RedBlackTreeNode)_setLeft(node *RedBlackTreeNode){
	if !n.IsNil {
		n.Left = node
	}else {
		panic("Children for nilNode is not allowed.")
	}
}

func (n *RedBlackTreeNode)_setRight(node *RedBlackTreeNode){
	if !n.IsNil{
		n.Right = node
		node.Parent = n
	}else {
		panic("Children for nilNode is not allowed.")
	}

}

