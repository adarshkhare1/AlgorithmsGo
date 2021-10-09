package searching

import (
	"Algorithms/testHelper"
	"testing"
)

func TestSortedArraySearch(t *testing.T) {
	input := []int{2, 4, 5, 7, 8, 13, 17, 20, 25}
	testHelper.VerifyIntsAreEqual(t, -1, FindIndexForItem(input, 12))
	testHelper.VerifyIntsAreEqual(t, 0, FindIndexForItem(input, 2))
	testHelper.VerifyIntsAreEqual(t, len(input)-1, FindIndexForItem(input, 25))
	testHelper.VerifyIntsAreEqual(t, 3, FindIndexForItem(input, 7))
	testHelper.VerifyIntsAreEqual(t, 6, FindIndexForItem(input, 17))
	testHelper.VerifyIntsAreEqual(t, -1, FindIndexForItem(input, 50))

}

func TestSortedArraySearchSingleElement(t *testing.T) {
	input := []int{12}
	testHelper.VerifyIntsAreEqual(t, 0, FindIndexForItem(input, 12))
	testHelper.VerifyIntsAreEqual(t, -1, FindIndexForItem(input, 1))
}
