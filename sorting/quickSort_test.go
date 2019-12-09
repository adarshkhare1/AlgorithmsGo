package sorting

import (
	"Algorithms/testHelper"
	"testing"
)
func TestQuicksort(t *testing.T) {
	a := []int {1, 4, 2, 5, 3}
	expected := []int {1, 2, 3, 4, 5}
	result := QuickSort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestQuickSortortEmpty(t *testing.T) {
	a := []int {}
	expected := []int {}
	result := QuickSort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestQuickSortPreSorted(t *testing.T) {
	a := []int {1, 2, 3, 4, 5}
	expected := []int {1, 2, 3, 4, 5}
	result := QuickSort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestQuickSortReverseSorted(t *testing.T) {
	a := []int {5, 4, 3, 2, 1}
	expected := []int {1, 2, 3, 4, 5}
	result := QuickSort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestQuickSortAllEqual(t *testing.T) {
	a := []int {5, 5, 5, 5, 5}
	expected := []int {5, 5, 5, 5, 5}
	result := QuickSort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}
