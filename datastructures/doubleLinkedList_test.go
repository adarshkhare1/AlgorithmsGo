package datastructures

import (
	"strconv"
	"testing"
)

func verifyDoubleLinkedList(t *testing.T, ll Doublelinkedlist, expected string){
	/* read the list until head is not nil */
	result := ""
	for cur := ll.head; cur != nil; cur = cur.next {
		result += " " + strconv.Itoa(cur.val)
	}
	verifyString(t, expected, result)
}

func TestDoublelinkedlist(t *testing.T) {
	ll := Doublelinkedlist{}

	ll.AddAtBegin(10)
	ll.AddAtEnd(20)
	verifyDoubleLinkedList(t, ll, " 10 20")
	ll.AddAtEnd(30)
	verifyDoubleLinkedList(t, ll, " 10 20 30")

	ll.DeleteAtBegin()
	verifyDoubleLinkedList(t, ll, " 20 30")

	ll.DeleteAtEnd()
	verifyDoubleLinkedList(t, ll, " 20")

	ll.DeleteAtBegin()
	verifyDoubleLinkedList(t, ll, "")

	ll.DeleteAtBegin()
	verifyDoubleLinkedList(t, ll, "")
}

func TestDoublelinkedlist_Reverse(t *testing.T) {
	ll := Doublelinkedlist{}
	ll.AddAtBegin(10)
	ll.AddAtEnd(20)
	ll.AddAtEnd(30)
	verifyDoubleLinkedList(t, ll, " 10 20 30")
	ll.Reverse()
	verifyDoubleLinkedList(t, ll, " 30 20 10")
}
