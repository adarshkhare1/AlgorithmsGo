package sorting

import (
	"Algorithms/testHelper"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	in := NewSortBasic()
	result := InsertionSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestInsertionSortEmpty(t *testing.T) {
	in := NewSortEmpty()
	result := InsertionSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestInsertionSortPreSorted(t *testing.T) {
	in := NewSortPreSorted()
	result := InsertionSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestInsertionSortReverseSorted(t *testing.T) {
	in := NewSortReverseSorted()
	result := InsertionSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestInsertionSortAllEqual(t *testing.T) {
	in := NewSortAllEquals()
	result := InsertionSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}
