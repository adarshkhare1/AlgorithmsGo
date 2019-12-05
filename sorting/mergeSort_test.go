package sorting

import (
	"Algorithms/testHelper"
	"testing"
)

func TestMergesort(t *testing.T) {
	a := []int {1, 4, 2, 5, 3}
	expected := []int {1, 2, 3, 4, 5}
	result := Mergesort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestMergesortEmpty(t *testing.T) {
	a := []int {}
	expected := []int {}
	result := Mergesort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestMergesortPreSorted(t *testing.T) {
	a := []int {1, 2, 3, 4, 5}
	expected := []int {1, 2, 3, 4, 5}
	result := Mergesort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestMergesortReverseSorted(t *testing.T) {
	a := []int {5, 4, 3, 2, 1}
	expected := []int {1, 2, 3, 4, 5}
	result := Mergesort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}

func TestMergesortAllEqual(t *testing.T) {
	a := []int {5, 5, 5, 5, 5}
	expected := []int {5, 5, 5, 5, 5}
	result := Mergesort(a)
	testHelper.VerifyArraysAreEqual(t, expected, result)
}
