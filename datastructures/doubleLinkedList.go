// demonstration of doubly linked list in golang
package datastructures

import "errors"

type doubleLinkedListNode struct {
	_value   int
	Next     *doubleLinkedListNode
	Previous *doubleLinkedListNode
}

type Doublelinkedlist struct {
	Head *doubleLinkedListNode
}

// to avoid mistakes when using pointer vs struct for new node creation
func newDoubleLinkedListNode(val int) *doubleLinkedListNode {
	n := &doubleLinkedListNode{}
	n._value = val
	n.Next = nil
	n.Previous = nil
	return n
}

func (ll *Doublelinkedlist) Append(val int) {
	n := newDoubleLinkedListNode(val)
	if ll.Head == nil {
		ll.Head = n
		return
	} else {
		cur := ll.Head
		for ; cur.Next != nil; cur = cur.Next {
		}
		cur.Next = n
		n.Previous = cur
	}
}

func (ll *Doublelinkedlist) InsertAtFront(val int) error {
	return ll.InsertAt(0, val)
}

func (ll *Doublelinkedlist) InsertAt(index uint, val int)  error {
	newNode := newDoubleLinkedListNode(val)
	if index == 0{
		newNode.Next = ll.Head
		ll.Head = newNode
		return nil
	} else {
		next := ll.Head.Next
		for i := uint(0); i < index-1; i++ {
			if next == nil {
				return errors.New("list is too small.")
			}
			next = next.Next
		}
		next.Previous.Next = newNode
		newNode.Previous = next.Previous
		next.Previous = newNode
		newNode.Next = next
	}
	return nil
}

func (ll *Doublelinkedlist) DeleteAt(index uint)  error {
	if index == 0 && ll.Head != nil{
		temp := ll.Head
		ll.Head = ll.Head.Next
		if ll.Head != nil {
			ll.Head.Previous = nil
		}
		temp.Next = nil
	} else {
		next := ll.Head
		for i := uint(0); i < index-1; i++ {
			if next == nil {
				return errors.New("list is too small.")
			}
			next = next.Next
		}
		if next != nil {
			if next.Previous != nil {
				next.Previous.Next = next.Next
			}
			if next.Next != nil {
				next.Next.Previous = next.Previous
				next.Next = next.Next.Next
			}
			next.Previous = nil
		} else {
			return errors.New("list is too small.")
		}
	}
	return nil

}

func (ll *Doublelinkedlist) DeleteAtEnd() error {
	// no item
	if ll.Head == nil {
		return errors.New("list is too small.")
	} else {
		cur := ll.Head
		for ; cur.Next.Next != nil; cur = cur.Next {
		}
		cur.Next.Previous = nil
		cur.Next = nil
	}
	return nil
}

func (ll *Doublelinkedlist) Count() int {
	var ctr int = 0

	for cur := ll.Head; cur != nil; cur = cur.Next {
		ctr += 1
	}

	return ctr
}

func (ll *Doublelinkedlist) Reverse() {
	var prev, next *doubleLinkedListNode
	cur := ll.Head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		cur.Previous = next
		prev = cur
		cur = next
	}

	ll.Head = prev
}
