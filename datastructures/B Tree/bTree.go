package B_Tree


type BTree struct {
	Root *BTreeNode
}

//NewBTree creates an empty BTree
func NewBTree() *BTree {
	node := NewBTreeNode()
	node.isLeaf = true
	diskWrite(node.fileName)
	return &BTree{Root: node}
}

//SearchTreeNode search a given key in BTree, return nil if key not found.
func SearchTreeNode (node *BTreeNode, key int) *BTreeNode {
	i :=0
	for i < len(node.items) && key > node.Key {
		i++
	}
	if i < len(node.items) && key == node.items[i].Key {
		return  &node.items[i]
	} else if node.isLeaf {
		return nil
	} else { //Load child node from the disk
		n := diskRead(node.fileName)
		return SearchTreeNode(n, key)
	}
}

