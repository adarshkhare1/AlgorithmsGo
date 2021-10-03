package datastructures

import (
	"Algorithms/testHelper"
	"testing"
)

func TestPriorityQueueBasic(t *testing.T) {
	q := buildTestPriorityQueue()
	items := make([]int, 0)
	items = append(items, q.Peak())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	testHelper.VerifyArraysAreEqual(t, []int{25, 25, 20, 17}, items)
}

func TestPriorityQueueUpdateMax(t *testing.T) {
	q := buildTestPriorityQueue()
	q.UpdateValue(2, 50)
	items := make([]int, 0)
	items = append(items, q.Peak())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	testHelper.VerifyArraysAreEqual(t, []int{50, 50, 25, 17, 13}, items)
}

func TestPriorityQueueUpdateMiddle(t *testing.T) {
	q := buildTestPriorityQueue()
	q.UpdateValue(2, 22)
	items := make([]int, 0)
	items = append(items, q.Peak())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	testHelper.VerifyArraysAreEqual(t, []int{25, 25, 22, 17, 13}, items)
}

func TestPriorityQueueRemove(t *testing.T) {
	q := buildTestPriorityQueue()
	q.Remove(3)
	items := make([]int, 0)
	items = append(items, q.Peak())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	items = append(items, q.Dequeue())
	testHelper.VerifyArraysAreEqual(t, []int{25, 25, 20, 17}, items)
}

func buildTestPriorityQueue() *PriorityQueue {
	q := NewPriorityQueue()
	q.EnqueueItems([]int{5, 13, 2, 25})
	q.EnqueueItems([]int{7, 17, 20, 8, 4})
	return q
}