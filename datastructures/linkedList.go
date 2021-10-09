package datastructures

import "errors"

/* value is the value of node; next is the pointer to next node */
type linkedListNode struct {
	_value int
	Next   *linkedListNode
}

type Linkedlist struct {
	Head *linkedListNode
}

func NewLinkedlist(head *linkedListNode) *Linkedlist {
	return &Linkedlist{Head: head}
}

func (l *Linkedlist) Append(val int)  {
	newNode := linkedListNode{_value: val, Next:   nil}
	if l.Head == nil {
		l.Head = &newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newNode
	}
}

func (l *Linkedlist) InsertAtFront(val int)  error{
	return l.InsertAt(0, val)
}

func (l *Linkedlist) InsertAt(index, val int)  error {
	newNode := linkedListNode{_value: val, Next: nil}
	if index == 0 {
		newNode.Next = l.Head
		l.Head = &newNode
	} else {
		node, err := l.getNodeByIndex(index-1)
		if err != nil {
			return err
		}
		next := node.Next
		node.Next = &newNode
		newNode.Next = next
	}
	return nil
}

func (l *Linkedlist) DeleteAt(index int)  error {
	if index == 0 && l.Head != nil{
		temp := l.Head.Next
		l.Head.Next = nil
		l.Head = temp
	} else {
		node, err := l.getNodeByIndex(index-1)
		if err != nil {
			return err
		}
		next := node.Next
		if next != nil {
			node.Next = next.Next
			next.Next = nil
		} else {
			return errors.New("list is too small.")
		}
	}
	return nil

}

func (l *Linkedlist) PopFrontValue() (int, error) {
	if l.Head == nil {
		return -1, errors.New("LinkedList is empty.")
	} else {
		temp := l.Head
		l.Head = l.Head.Next
		return temp._value, nil
	}
}

func (l *Linkedlist) PopBackValue() (int, error) {
	if l.Head == nil {
		return -1, errors.New("LinkedList is empty.")
	} else {
		current := l.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
		return current._value, nil
	}
}

func (l *Linkedlist) getNodeByIndex(index int) (*linkedListNode, error) {
	current := l.Head
	next := current.Next
	for i := 0; i < index; i++ {
		current = next
		if current == nil {
			return nil, errors.New("list is too small.")
		}
		next = next.Next
	}
	return current, nil
}