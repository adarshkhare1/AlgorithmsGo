package datastructures

import (
	"testing"
)

func TestStack_Push(t *testing.T) {

	s := NewStack(10)
	err := s.Push(2)
	verifyInt(t, 1, s.Count())
	verifyErrorNil(t, err)
}

func TestStack_Peek(t *testing.T) {
	s := NewStack(10)
	s.Push(2)
	n, err := s.Peek()
	verifyInt(t, 2, n)
	verifyInt(t, 1, s.Count())
	verifyErrorNil(t, err)

}

func TestStack_Pop(t *testing.T) {
	s := NewStack(10)
	s.Push(2)
	n, err := s.Pop()
	verifyInt(t, 2, n)
	verifyInt(t, 0, s.Count())
	verifyErrorNil(t, err)
}

func TestStack_Empty(t *testing.T) {
	s := NewStack(10)
	n,err := s.Pop()
	verifyInt(t, -1, n)
	verifyErrorNotNil(t, err)
}

func TestStack_Overflow(t *testing.T) {
	capacity := 3
	s := NewStack(capacity)

	for ;capacity > 0; capacity-- {
		err := s.Push(capacity)
		verifyErrorNil(t, err)
	}
	err := s.Push(10)
	verifyErrorNotNil(t, err)
	n, err := s.Pop()
	verifyInt(t, 1, n)
	verifyInt(t, 2, s.Count())
}