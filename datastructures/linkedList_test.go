package datastructures

import (
	"Algorithms/testHelper"
	"strconv"
	"testing"
)

func verifyLinkedList(t *testing.T, list *Linkedlist, expected string){
	/* read the list until head is not nil */
	result := ""
	n := list.Head
	for n != nil {
		result += " " + strconv.Itoa(n._value)
		n = n.Next /*head points to next node */
	}
	testHelper.VerifyStringsAreEqual(t, expected, result)
}


func TestNode(t *testing.T) {
	lista := NewLinkedlist(nil)
	lista.Append(5)
	lista.Append(14)
	lista.Append(12)
	verifyLinkedList(t, lista, " 5 14 12")

	lista.InsertAt(1,56)
	verifyLinkedList(t, lista, " 5 56 14 12")

	lista.InsertAtFront(36)
	verifyLinkedList(t, lista, " 36 5 56 14 12")

	lista.PopFrontValue()
	verifyLinkedList(t, lista, " 5 56 14 12")

	lista.PopBackValue()
	verifyLinkedList(t, lista, " 5 56 14")

	lista.DeleteAt(1)
	verifyLinkedList(t, lista, " 5 14")

}