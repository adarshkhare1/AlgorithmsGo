package datastructures

import (
	"errors"
)

type Stack struct {
	value []int
	_top  int
}

func NewStack(capacity int) *Stack {
	arr := make([]int, capacity)
	return &Stack{value: arr}
}


func (s *Stack) Pop() (int, error){
	if s._top > 0 {
		s._top--
		return s.value[s._top], nil
	} else {
		return -1, errors.New("Stack is empty")
	}

}

func (s *Stack) Push(val int) error{
	if s._top < len(s.value) {
		s.value[s._top] = val
		s._top++
		return nil
	} else {
		return errors.New("Stack is full")
	}
}

func (s *Stack) Peek() (int, error){
	if s._top > 0 {
		return s.value[s._top-1], nil
	} else {
		return -1, errors.New("Stack is empty")
	}
}

func (s *Stack) Count() int {
	return s._top
}