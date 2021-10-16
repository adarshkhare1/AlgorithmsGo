package redBlackTree

import (
	"strconv"
)

type RedBlackTree struct {
	Root *RedBlackTreeNode
}



func NewRedBlackBinaryTreeEmpty() *RedBlackTree {
	return &RedBlackTree{Root: _nilRedBlackTreeNode(nil)}
}

func NewRedBlackTree(rootValue int) *RedBlackTree {
	root := _newRedBlackTreeNode(rootValue)
	return &RedBlackTree{Root: root}
}

func (t *RedBlackTree) Depth() uint {
	return _calculateRBTreeDepth(t.Root, 0)
}

func (t *RedBlackTree) Insert(value int){
	z := _newRedBlackTreeNode(value)
	x := t.Root
	y := _nilRedBlackTreeNode(x)
	for !x.IsNil{
		y = x
		if z.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if y.IsNil{
		t.Root = z
	} else  if z.Value < y.Value {
		y._setLeft(z)
	} else {
		y._setRight(z)
	}
	z.Color = Red
	t.insertFixup(z)
}

func (t *RedBlackTree) Delete(z *RedBlackTreeNode){
	if z.IsNil{
		return // nil node cannot be deleted from the tree
	}
	y := z
	yColor := y.Color
	var x *RedBlackTreeNode = nil
	if z.Left.IsNil{
		x = z.Right
		t.transplant(z, z.Right)
	}else if z.Right.IsNil{
		x = z.Left
		t.transplant(z, z.Left)
	}else {
		y = z.Right.TreeMinimum()
		yColor = y.Color
		if y.Parent == z{
			if x != nil {
				x.Parent = y
			}
		}else {
			t.transplant(y, y.Right)
			y._setRight(z.Right)
		}
		t.transplant(z, y)
		y._setLeft(z.Left)
		y.Color = z.Color
	}
	if yColor == Black {
		t.deleteFixup(x)
	}
}

func (n *RedBlackTreeNode) TreeMinimum() *RedBlackTreeNode {
	node := n
	for !node.Left.IsNil{
		node = node.Left
	}
	return node
}

func (n *RedBlackTreeNode) TreeMaximum() *RedBlackTreeNode {
	node := n
	for !node.Right.IsNil{
		node = node.Right
	}
	return  node
}

func (t *RedBlackTree)Inorder() string{
	return _rbInorder(t.Root, "")
}

// SearchTreeNode return the node that has value equal to given value.
//Return nil if no matching value found.
func SearchTreeNode (root *RedBlackTreeNode, value int) *RedBlackTreeNode {
	if root == nil || root.Value == value {
		return root
	}
	if value < root.Value{
		return SearchTreeNode(root.Left, value)
	}else{
		return SearchTreeNode(root.Right, value)
	}
}

// Inorder traversal print left node, Root, right node order
func _rbInorder(n *RedBlackTreeNode, result string) string {
	if n.IsNil {
		return result
	}
	if !n.Left.IsNil {
		result = _rbInorder(n.Left, result)
	}
	result += " " + strconv.Itoa(n.Value)
	if !n.Right.IsNil {
		result = _rbInorder(n.Right, result)
	}
	return result
}

func (t *RedBlackTree) leftRotate(x *RedBlackTreeNode){
	if x != nil && !x.IsNil {
		y := x.Right
		x._setRight(y.Left)
		y.Parent = x.Parent
		if x.Parent == nil {
			t.Root = y
		} else if x == x.Parent.Left {
			x.Parent._setLeft(y)
		} else {
			x.Parent._setRight(y)
		}
		y._setLeft(x)
	}
}

func (t *RedBlackTree) rightRotate(y *RedBlackTreeNode){
	if y != nil && !y.IsNil {
		x := y.Left
		y._setLeft(x.Right)
		x.Parent = y.Parent
		if y.Parent == nil {
			t.Root = x
		} else if y == y.Parent.Right {
			x.Parent._setRight(y)
		}else {
			x.Parent._setLeft(y)
		}
		x._setRight(y)
	}
}

func (t *RedBlackTree) insertFixup(z *RedBlackTreeNode) {
	if z.Parent == nil || z.Parent.Parent == nil{
		return // it is root, nothing to fixup
	}
	for z.Parent.Color == Red {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == Red {
				z.Parent.Color = Black
				y.Color = Black
				z.Parent.Parent.Color = Red
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.leftRotate(z)
				}
				z.Parent.Color = Black
				z.Parent.Parent.Color = Red
				t.rightRotate(z.Parent.Parent)
			}
		} else {
			if z.Parent == z.Parent.Parent.Right {
				y := z.Parent.Parent.Left
				if y.Color == Red {
					z.Parent.Color = Black
					y.Color = Black
					z.Parent.Parent.Color = Red
					z = z.Parent.Parent
				} else if z == z.Parent.Left {
					z = z.Parent
					t.rightRotate(z)
				}
				if z.Parent != nil {
					z.Parent.Color = Black
					if z.Parent.Parent != nil {
						z.Parent.Parent.Color = Red
					}
				}
				if z.Parent != nil && z.Parent.Parent != nil {
					t.leftRotate(z.Parent.Parent)
				}
			}
		}
		if z.Parent == nil || z.Parent.Parent == nil{
			break
		}
	}
	t.Root.Color = Black
}

func (t *RedBlackTree) deleteFixup(x *RedBlackTreeNode){
	if x.Parent == nil{
		return // it is root, nothing to fixup
	}
	for x != t.Root && x.Color == Black {
		if x == x.Parent.Left{
			w := x.Parent.Right
			if w.Color == Red {
				w.Color = Black
				x.Parent.Color = Red
				t.leftRotate(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == Black && w.Right.Color == Black {
				w.Color = Red
				x = x.Parent
			} else {
				if w.Right.Color == Black {
					w.Left.Color = Black
					w.Color = Red
					t.rightRotate(w)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = Black
				x.Right.Color = Black
				t.leftRotate(x.Parent)
				x = t.Root
			}
		}else {
			w := x.Parent.Left
			if w.Color == Red {
				w.Color = Black
				x.Parent.Color = Red
				t.rightRotate(x.Parent)
				w = x.Parent.Left
			}
			if w.Left.Color == Black && w.Right.Color == Black {
				w.Color = Red
				x = x.Parent
			} else {
				if w.Left.Color == Black {
					w.Right.Color = Black
					w.Color = Red
					t.leftRotate(w)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = Black
				x.Left.Color = Black
				t.rightRotate(x.Parent)
				x = t.Root
			}
		}
	}
	x.Color = Black
}

func (t *RedBlackTree) transplant(u *RedBlackTreeNode, v *RedBlackTreeNode){
	if u.Parent.IsNil{
		t.Root = v
	}else if u == u.Parent.Left{
		u.Parent._setLeft(v)
	}else {
		u.Parent._setRight(v)
	}
}

// helper function for t.depth
func _calculateRBTreeDepth(n *RedBlackTreeNode, depth uint) uint {
	if n.IsNil {
		return depth
	}
	return _max(_calculateRBTreeDepth(n.Left, depth+1), _calculateRBTreeDepth(n.Right, depth+1))
}

func _max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}