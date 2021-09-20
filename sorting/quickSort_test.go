package sorting

import (
	"Algorithms/testHelper"
	"testing"
)
func TestQuicksort(t *testing.T) {
	in := NewSortBasic()
	result := QuickSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestQuickSortEmpty(t *testing.T) {
	in := NewSortEmpty()
	result := QuickSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestQuickSortPreSorted(t *testing.T) {
	in := NewSortPreSorted()
	result := QuickSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestQuickSortReverseSorted(t *testing.T) {
	in := NewSortReverseSorted()
	result := QuickSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}

func TestQuickSortAllEqual(t *testing.T) {
	in := NewSortAllEquals()
	result := QuickSort(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}
