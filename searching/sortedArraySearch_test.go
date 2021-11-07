package searching

import (
	"Algorithms/testHelper"
	"testing"
)


func TestSortedArraySearch(t *testing.T) {
	input := getTestList()
	testHelper.VerifyIntsAreEqual(t, -1, FindItemIndexSequentialSkipSearch(input, 12))
	testHelper.VerifyIntsAreEqual(t, 0, FindItemIndexSequentialSkipSearch(input, 2))
	testHelper.VerifyIntsAreEqual(t, len(input)-1, FindItemIndexSequentialSkipSearch(input, 25))
	testHelper.VerifyIntsAreEqual(t, 3, FindItemIndexSequentialSkipSearch(input, 7))
	testHelper.VerifyIntsAreEqual(t, 6, FindItemIndexSequentialSkipSearch(input, 17))
	testHelper.VerifyIntsAreEqual(t, -1, FindItemIndexSequentialSkipSearch(input, 50))

}

func TestSortedArraySearchSingleElement(t *testing.T) {
	input := []int{12}
	testHelper.VerifyIntsAreEqual(t, 0, FindItemIndexSequentialSkipSearch(input, 12))
	testHelper.VerifyIntsAreEqual(t, -1, FindItemIndexSequentialSkipSearch(input, 1))
}

func TestSortedArrayBinarySearch(t *testing.T) {
	input := getTestList()
	testHelper.VerifyIntsAreEqual(t, -1, FindItemIndexBinarySearch(input, 12))
	testHelper.VerifyIntsAreEqual(t, 0, FindItemIndexBinarySearch(input, 2))
	testHelper.VerifyIntsAreEqual(t, len(input)-1, FindItemIndexBinarySearch(input, 25))
	testHelper.VerifyIntsAreEqual(t, 3, FindItemIndexBinarySearch(input, 7))
	testHelper.VerifyIntsAreEqual(t, 6, FindItemIndexBinarySearch(input, 17))
	testHelper.VerifyIntsAreEqual(t, -1, FindItemIndexBinarySearch(input, 50))

}

func TestSortedArrayBinarySearchSingleElement(t *testing.T) {
	input := []int{12}
	testHelper.VerifyIntsAreEqual(t, 0, FindItemIndexBinarySearch(input, 12))
	testHelper.VerifyIntsAreEqual(t, -1, FindItemIndexBinarySearch(input, 1))
}

func getTestList() []int {
	return []int{2, 4, 5, 7, 8, 13, 17, 20, 25}
}