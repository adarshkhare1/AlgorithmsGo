package datastructures

import (
	"strconv"
	"testing"
)

func verifyLinkedList(t *testing.T, head *Node, expected string){
	/* read the list until head is not nil */
	result := ""
	for head != nil {
		result += " " + strconv.Itoa(head.Value)
		head = head.Next /*head points to next node */
	}
	verifyString(t, expected, result)
}


func TestNode(t *testing.T) {
	lista := new(Node)
	lista.PushBack(5).PushBack(14).PushBack(12)
	verifyLinkedList(t, lista, " 5 14 12")

	lista.PushBack(56)
	verifyLinkedList(t, lista, " 5 14 12 56")

	lista = lista.PushFront(36)
	verifyLinkedList(t, lista, " 36 5 14 12 56")

	lista = lista.PopFront()
	verifyLinkedList(t, lista, " 5 14 12 56")

	lista.PopBack()
	verifyLinkedList(t, lista, " 5 14 12")

}