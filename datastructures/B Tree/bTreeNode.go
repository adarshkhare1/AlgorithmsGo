package B_Tree

type BTreeNode struct {
	Key int
	items []BTreeNode
	fileName string
	isLeaf bool
}

//NewBTreeNode creates an empty BTreeNode
func NewBTreeNode() *BTreeNode {
	return &BTreeNode{}
}

