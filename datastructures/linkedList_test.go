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


func TestLinkedListOperations(t *testing.T) {
	lista := NewLinkedlist(nil)
	lista.Append(5)
	lista.Append(14)
	lista.Append(12)
	verifyLinkedList(t, lista, " 5 14 12")

	lista.InsertAt(1, 56)
	verifyLinkedList(t, lista, " 5 56 14 12")

	lista.InsertAtFront(36)
	verifyLinkedList(t, lista, " 36 5 56 14 12")

	lista.PopFrontValue()
	verifyLinkedList(t, lista, " 5 56 14 12")

	lista.PopBackValue()
	verifyLinkedList(t, lista, " 5 56 14")

}

func TestEmptyListErrors(t *testing.T) {
	lista := NewLinkedlist(nil)
	_, err := lista.PopBackValue()
	testHelper.VerifyErrorNotNil(t, err)
	_, err = lista.PopFrontValue()
	testHelper.VerifyErrorNotNil(t, err)
}

func TestLinkedlist_InsertDelete(t *testing.T) {
	lista := NewLinkedlist(nil)
	lista.Append(5)
	lista.Append(14)
	lista.Append(12)
	verifyLinkedList(t, lista, " 5 14 12")

	lista.InsertAt(1,56)
	verifyLinkedList(t, lista, " 5 56 14 12")

	lista.DeleteAt(0)
	verifyLinkedList(t, lista, " 56 14 12")

	lista.DeleteAt(1)
	verifyLinkedList(t, lista, " 56 12")

	lista.DeleteAt(1)
	verifyLinkedList(t, lista, " 56")

	err := lista.DeleteAt(1)
	testHelper.VerifyErrorNotNil(t, err)

	err = lista.InsertAt(1, 5)
	verifyLinkedList(t, lista, " 56 5")
	testHelper.VerifyErrorNil(t, err)

	err = lista.InsertAt(3, 5)
	testHelper.VerifyErrorNotNil(t, err)

}