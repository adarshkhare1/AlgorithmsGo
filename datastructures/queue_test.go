package datastructures

import "testing"

func TestQueue_Operation(t *testing.T) {
	capacity := 5
	q := NewQueue(capacity)
	_fillQueue(q, t)
	_emptyQueue(q, t)
}

func TestQueue_Overflow(t *testing.T) {
	capacity := 5
	q := NewQueue(capacity)
	_fillQueue(q, t)
	err := q.Enqueue(capacity)
	verifyErrorNotNil(t, err)
}

func TestQueue_Empty(t *testing.T) {
	capacity := 3
	q := NewQueue(capacity)
	_fillQueue(q, t)
	_emptyQueue(q, t)
	n, err := q.Dequeu()
	verifyInt(t, -1, n)
	verifyErrorNotNil(t, err)
}

func TestQueue_Rotation(t *testing.T) {
	capacity := 5
	q := NewQueue(capacity)
	_fillQueue(q, t)
	// Take few out
	for i := 0; i < capacity - 2; i++ {
		n, err := q.Dequeu()
		verifyInt(t, i, n)
		verifyErrorNil(t, err)
	}
	// Add couple of more to trigger rotation of array
	err := q.Enqueue(1)
	verifyErrorNil(t, err)
	err = q.Enqueue(2)
	verifyErrorNil(t, err)

	//Take 2 more out to hit bottom of the array
	n := 0
	n, err = q.Dequeu()
	verifyInt(t, capacity-2, n)
	verifyErrorNil(t, err)

	n, err = q.Dequeu()
	verifyInt(t, capacity-1, n)
	verifyErrorNil(t, err)

	//Take one more out to trigger dequeue rotation
	n, err = q.Dequeu()
	verifyInt(t, 1, n)
	verifyErrorNil(t, err)
}

func _fillQueue(q *Queue, t *testing.T) {
	capacity := q.Capacity()
	for i := 0; i < capacity; i++ {
		err := q.Enqueue(i)
		verifyErrorNil(t, err)
	}
}

func _emptyQueue(q *Queue, t *testing.T) {
	capacity := q.Capacity()
	for i := 0; i < capacity; i++ {
		n, err := q.Dequeu()
		verifyInt(t, i, n)
		verifyErrorNil(t, err)
	}
}