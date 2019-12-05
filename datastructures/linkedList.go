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
		previous := l.Head
		next := previous.Next
		for i := 0; i < index-1; i++ {
			previous = next
			if previous == nil {
				return errors.New("list is too small.")
			}
			next = next.Next
		}
		previous.Next = &newNode
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
		previous := l.Head
		next := previous.Next
		for i := 0; i < index-1; i++ {
			previous = next
			if previous == nil {
				return errors.New("list is too small.")
			}
			next = next.Next
		}
		previous.Next = next.Next
		next.Next = nil
	}
	return nil

}

func (l *Linkedlist) PopFrontValue() (int, error) {
	if l.Head == nil {
		return -1, errors.New("LinkedList is empty.")
	} else {
		temp := l.Head
		/* now head is equals cpnode (second node) */
		l.Head = l.Head.Next
		return temp._value, nil
	}
}

func (l *Linkedlist) PopBackValue() (int, error) {
	if l.Head == nil {
		return -1, errors.New("LinkedList is empty.")
	} else {
		current := l.Head
		/* read list to the penultimate node */
		for current.Next.Next != nil {
			current = current.Next
		}
		/* the penultimate node points to null. In this way the last node is deleted */
		current.Next = nil
		return current._value, nil
	}
}

