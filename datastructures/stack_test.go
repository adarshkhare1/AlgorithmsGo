package datastructures

import (
	"AlgorithmsGo/testHelper"
	"testing"
)

func TestStack_Push(t *testing.T) {
	
	s := NewStack(10)
	err := s.Push(2)
	testHelper.VerifyIntsAreEqual(t, 1, s.Count())
	testHelper.VerifyErrorNil(t, err)
}

func TestStack_Peek(t *testing.T) {
	s := NewStack(10)
	s.Push(2)
	n, err := s.Peek()
	testHelper.VerifyIntsAreEqual(t, 2, n)
	testHelper.VerifyIntsAreEqual(t, 1, s.Count())
	testHelper.VerifyErrorNil(t, err)
	
}

func TestStack_Pop(t *testing.T) {
	s := NewStack(10)
	s.Push(2)
	n, err := s.Pop()
	testHelper.VerifyIntsAreEqual(t, 2, n)
	testHelper.VerifyIntsAreEqual(t, 0, s.Count())
	testHelper.VerifyErrorNil(t, err)
}

func TestStack_Empty(t *testing.T) {
	s := NewStack(10)
	n, err := s.Pop()
	testHelper.VerifyIntsAreEqual(t, -1, n)
	testHelper.VerifyErrorNotNil(t, err)
}

func TestStack_Overflow(t *testing.T) {
	capacity := 3
	s := NewStack(capacity)
	
	for ; capacity > 0; capacity-- {
		err := s.Push(capacity)
		testHelper.VerifyErrorNil(t, err)
	}
	err := s.Push(10)
	testHelper.VerifyErrorNotNil(t, err)
	n, err := s.Pop()
	testHelper.VerifyIntsAreEqual(t, 1, n)
	testHelper.VerifyIntsAreEqual(t, 2, s.Count())
}
