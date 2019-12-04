// demonstration of doubly linked list in golang
package datastructures

type doubleLinkedListNode struct {
	val  int
	next *doubleLinkedListNode
	prev *doubleLinkedListNode
}

type Doublelinkedlist struct {
	head *doubleLinkedListNode
}

// to avoid mistakes when using pointer vs struct for new node creation
func newDoubleLinkedListNode(val int) *doubleLinkedListNode {
	n := &doubleLinkedListNode{}
	n.val = val
	n.next = nil
	n.prev = nil
	return n
}

func (ll *Doublelinkedlist) AddAtBegin(val int) {
	n := newDoubleLinkedListNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *Doublelinkedlist) AddAtEnd(val int) {
	n := newDoubleLinkedListNode(val)

	if ll.head == nil {
		ll.head = n
		return
	}

	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n
	n.prev = cur
}

func (ll *Doublelinkedlist) DeleteAtBegin() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	ll.head = cur.next

	if ll.head != nil {
		ll.head.prev = nil
	}

	return cur.val
}

func (ll *Doublelinkedlist) DeleteAtEnd() int {
	// no item
	if ll.head == nil {
		return -1
	}

	// only one item
	if ll.head.next == nil {
		return ll.DeleteAtBegin()
	}

	// more than one, go to second last
	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {
	}

	retval := cur.next.val
	cur.next = nil
	return retval
}

func (ll *Doublelinkedlist) Count() int {
	var ctr int = 0

	for cur := ll.head; cur != nil; cur = cur.next {
		ctr += 1
	}

	return ctr
}

func (ll *Doublelinkedlist) Reverse() {
	var prev, next *doubleLinkedListNode
	cur := ll.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}

	ll.head = prev
}
