package sorting

import (
	"Algorithms/testHelper"
	"testing"
)

func TestMergesort(t *testing.T) {
	in := NewSortBasic()
	result := Mergesort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestMergesortEmpty(t *testing.T) {
	in := NewSortEmpty()
	result := Mergesort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestMergesortPreSorted(t *testing.T) {
	in := NewSortPreSorted()
	result := Mergesort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestMergesortReverseSorted(t *testing.T) {
	in := NewSortReverseSorted()
	result := Mergesort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestMergesortAllEqual(t *testing.T) {
	in := NewSortAllEquals()
	result := Mergesort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}
