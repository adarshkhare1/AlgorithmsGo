package datastructures

import "errors"

type Queue struct {
	value   []int
	_top    int
	_bottom int
}
func NewQueue(capacity int) *Queue {
	arr := make([]int, capacity)
	return &Queue{value: arr}
}
func (q *Queue) Enqueue(val int) error {
	if q.Length() == len(q.value) {
		return errors.New("Queue is full capacity.")
	} else {
		if q._top == len(q.value) {
			q._top = 0
		}
		q.value[q._top] = val
		q._top++
		return nil
	}
}

func (q * Queue) Dequeu() (int, error)  {
	if q._top == q._bottom {
		return -1, errors.New("Queue is empty.")
	} else {
		if q._bottom == len(q.value) {
			q._bottom = 0
		}
		q._bottom++
		return q.value[q._bottom-1], nil
	}
}

func (q *Queue) Capacity() int {
	return len(q.value)
}

func (q *Queue) Length() int  {
	if q._top > q._bottom {
		return q._top - q._bottom
	} else {
		return q._bottom - q._top
	}
}