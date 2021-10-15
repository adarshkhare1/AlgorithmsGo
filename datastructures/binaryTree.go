package datastructures

// package main

import (
	"strconv"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
	Parent *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTreeEmpty() *BinaryTree {
	return &BinaryTree{Root: nil}
}

func NewBinaryTree(rootValue int) *BinaryTree {
	root := _newTreeNode(rootValue)
	return &BinaryTree{Root: root}
}

func (t *BinaryTree) Depth() uint {
	return _calculateDepth(t.Root, 0)
}

func (t *BinaryTree)Insert(k int) *TreeNode{
	newNode := _newTreeNode(k)
	x := t.Root
	var y *TreeNode
	for x != nil {
		y = x
		if newNode.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	newNode.Parent = y
	if y == nil{ // Empty tree
		t.Root = newNode
	} else if newNode.Value < y.Value {
		y._setLeft(newNode)
	} else {
		y._setRight(newNode)
	}
	return newNode
}

func (t *BinaryTree)Delete (z *TreeNode){
	if z.Left == nil{
		t.transplant(z, z.Right)
	} else if z.Right == nil{
		t.transplant(z, z.Left)
	} else{
		y := treeMinimum(z.Right)
		if y.Parent != z{
			t.transplant(y, y.Right)
		}
		t.transplant(z, y)
		y._setLeft(z.Left)
	}
}

func SearchTreeNode (x *TreeNode, k int) *TreeNode{
	if x == nil || x.Value == k{
		return x
	}
	if k < x.Value{
		return SearchTreeNode(x.Left, k)
	}else{
		return SearchTreeNode(x.Right, k)
	}
}

func (t *BinaryTree)transplant(u *TreeNode, v *TreeNode){
	if u.Parent == nil{
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent._setLeft(v)
	} else if u.Parent.Right == u {
		u.Parent._setRight(v)
	}
}


func (t *BinaryTree)Inorder() string{
	return _inorder(t.Root, "")
}

func (t *BinaryTree)Preorder() string{
	return _preorder(t.Root, "")
}

func (t *BinaryTree)Postorder() string{
	return _postOrder(t.Root, "")
}

func _newTreeNode(val int) *TreeNode {
	n := &TreeNode{
		Value:val,
		Left: nil,
		Right:nil,
		Parent:nil,
	}
	return n
}

func (n *TreeNode)_setLeft(node *TreeNode){
	n.Left = node
	if node != nil {
		node.Parent = n
	}
}

func (n *TreeNode)_setRight(node *TreeNode){
	n.Right = node
	if node != nil {
		node.Parent = n
	}
}

// Inorder traversal print left node, Root, right node order
func _inorder(n *TreeNode, result string) string {
	if n == nil {
		return result
	}
	if n.Left != nil {
		result = _inorder(n.Left, result)
	}
	result += " " + strconv.Itoa(n.Value)
	if n.Right != nil {
		result = _inorder(n.Right, result)
	}
	return result
}

// PostOrder traversal print left node, right node, Root order
func _postOrder( n *TreeNode, result string) string {
	if n == nil {
		return result
	}
	if n.Left != nil {
		result = _postOrder(n.Left, result)
	}
	if n.Right != nil {
		result = _postOrder(n.Right, result)
	}
	result += " " + strconv.Itoa(n.Value)
	return result
}

// PreOrder traversal prints Root, left node, right node order
func _preorder(n *TreeNode, result string) string {
	if n == nil {
		return result
	}
	result += " " + strconv.Itoa(n.Value)
	if n.Left != nil {
		result = _preorder(n.Left, result)
	}
	if n.Right != nil {
		result = _preorder(n.Right, result)
	}
	return result
}


func treeMinimum(node *TreeNode) *TreeNode {
	for node.Left != nil{
		node = node.Left
	}
	return  node
}

func treeMaximum(node *TreeNode) *TreeNode {
	for node.Right != nil{
		node = node.Right
	}
	return  node
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


