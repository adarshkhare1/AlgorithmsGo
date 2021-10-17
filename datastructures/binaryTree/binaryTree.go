package binaryTree

// package main

import (
	"strconv"
)


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

type TreeNode struct {
	Key  int
	Left *TreeNode
	Right *TreeNode
	Parent *TreeNode
}


func (t *BinaryTree) Depth() uint {
	return _calculateDepth(t.Root, 0)
}

// Insert the new node with given value.
func (t *BinaryTree)Insert(value int) *TreeNode {
	newNode := _newTreeNode(value)
	x := t.Root
	var y *TreeNode
	for x != nil {
		y = x
		if newNode.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	newNode.Parent = y
	if y == nil{ // Empty tree
		t.Root = newNode
	} else if newNode.Key < y.Key {
		y._setLeft(newNode)
	} else {
		y._setRight(newNode)
	}
	return newNode
}

// Delete the given node from the tree.
func (t *BinaryTree)Delete (node *TreeNode){
	if node.Left == nil{
		t.transplant(node, node.Right)
	} else if node.Right == nil{
		t.transplant(node, node.Left)
	} else{
		y := treeMinimum(node.Right)
		if y.Parent != node {
			t.transplant(y, y.Right)
		}
		t.transplant(node, y)
		y._setLeft(node.Left)
	}
}

// SearchTreeNode return the node that has value equal to given value.
//  Return nil if no matching value found.
func SearchTreeNode (root *TreeNode, value int) *TreeNode {
	if root == nil || root.Key == value {
		return root
	}
	if value < root.Key {
		return SearchTreeNode(root.Left, value)
	}else{
		return SearchTreeNode(root.Right, value)
	}
}
// transplant replaces a subtree rooted at node with replacementNode.
func (t *BinaryTree)transplant(node *TreeNode, replacementNode *TreeNode){
	if node.Parent == nil{
		t.Root = replacementNode
	} else if node == node.Parent.Left {
		node.Parent._setLeft(replacementNode)
	} else if node.Parent.Right == node {
		node.Parent._setRight(replacementNode)
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

func _newTreeNode(k int) *TreeNode {
	n := &TreeNode{
		Key:    k,
		Left:   nil,
		Right:  nil,
		Parent: nil,
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
	result += " " + strconv.Itoa(n.Key)
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
	result += " " + strconv.Itoa(n.Key)
	return result
}

// PreOrder traversal prints Root, left node, right node order
func _preorder(n *TreeNode, result string) string {
	if n == nil {
		return result
	}
	result += " " + strconv.Itoa(n.Key)
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


