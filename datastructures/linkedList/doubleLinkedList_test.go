package linkedList

import (
	"AlgorithmsGo/testHelper"
	"strconv"
	"testing"
)

func verifyDoubleLinkedList(t *testing.T, ll Doublelinkedlist, expected string) {
	/* read the list until head is not nil */
	result := ""
	for cur := ll.Head; cur != nil; cur = cur.Next {
		result += " " + strconv.Itoa(cur._value)
	}
	testHelper.VerifyStringsAreEqual(t, expected, result)
}

func TestDoublelinkedlistOperations(t *testing.T) {
	ll := Doublelinkedlist{}
	
	ll.InsertAtFront(10)
	ll.Append(20)
	verifyDoubleLinkedList(t, ll, " 10 20")
	ll.Append(30)
	verifyDoubleLinkedList(t, ll, " 10 20 30")
	ll.InsertAt(1, 40)
	verifyDoubleLinkedList(t, ll, " 10 40 20 30")
}

func TestDoubleLinkedlist_InsertDelete(t *testing.T) {
	ll := Doublelinkedlist{}
	capacity := 5
	for i := 0; i < capacity; i++ {
		ll.Append(10 * (i + 1))
	}
	verifyDoubleLinkedList(t, ll, " 10 20 30 40 50")
	ll.DeleteAt(1)
	verifyDoubleLinkedList(t, ll, " 10 30 40 50")
	ll.DeleteAt(0)
	verifyDoubleLinkedList(t, ll, " 30 40 50")
	ll.DeleteAtEnd()
	verifyDoubleLinkedList(t, ll, " 30 40")
	err := ll.DeleteAt(1)
	verifyDoubleLinkedList(t, ll, " 30")
	testHelper.VerifyErrorNil(t, err)
	ll.DeleteAt(0)
	verifyDoubleLinkedList(t, ll, "")
	err = ll.DeleteAt(0)
	testHelper.VerifyErrorNotNil(t, err)
}

func TestDoublelinkedlist_Reverse(t *testing.T) {
	ll := Doublelinkedlist{}
	ll.InsertAtFront(10)
	ll.Append(20)
	ll.Append(30)
	verifyDoubleLinkedList(t, ll, " 10 20 30")
	ll.Reverse()
	verifyDoubleLinkedList(t, ll, " 30 20 10")
}
