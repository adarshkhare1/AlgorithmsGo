package datastructures

/* value is the value of node; next is the pointer to next node */
type Node struct {
	Value int
	Next  *Node
}

/* first node, called head. It points from first node to last node */
var head *Node = nil

func (n *Node) PushFront(val int) *Node {
	/* if there's no nodes, head points to l (first node) */
	if head == nil {
		n.Value = val
		n.Next = nil
		head = n
		return n
	} else {
		/* create a new node equals to head */
		nnode := new(Node)
		nnode = head
		/* create a second node with new value and `next -> nnode`
		*  is this way, nnode2 is before nnode
		 */
		nnode2 := &Node{
			Value: val,
			Next:  nnode,
		}
		/* now head is equals nnode2 */
		head = nnode2
		return head
	}
}

func (n *Node) PushBack(val int) *Node {
	/* if there's no nodes, head points to l (first node) */
	if head == nil {
		n.Value = val
		n.Next = nil
		head = n
		return n
	} else {
		/* read all list to last node */
		for n.Next != nil {
			n = n.Next
		}
		/* allocate a new portion of memory */
		n.Next = new(Node)
		n.Next.Value = val
		n.Next.Next = nil
		return n
	}
}

func (n *Node) PopFront() *Node {
	if head == nil {
		return head
	}
	/* create a new node equals to first node pointed by head */
	cpnode := new(Node)
	cpnode = head.Next

	/* now head is equals cpnode (second node) */
	head = cpnode

	return head
}

func (n *Node) PopBack() *Node {
	if head == nil {
		return head
	}
	/* create a new node equals to head */
	cpnode := new(Node)
	cpnode = head

	/* read list to the penultimate node */
	for cpnode.Next.Next != nil {
		cpnode = cpnode.Next
	}
	/* the penultimate node points to null. In this way the last node is deleted */
	cpnode.Next = nil
	return head
}

