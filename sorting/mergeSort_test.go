package sorting

import (
	"testing"
)

func TestMergesort(t *testing.T) {
	TestSort(MergeSort, NewSortBasic(), t)
}

func TestMergeSortSequence(t *testing.T) {
	TestSort(HeapSort, NewSortSequence(), t)
}

func TestMergesortEmpty(t *testing.T) {
	TestSort(MergeSort, NewSortEmpty(), t)
}

func TestMergesortPreSorted(t *testing.T) {
	TestSort(MergeSort, NewSortPreSorted(), t)
}

func TestMergesortReverseSorted(t *testing.T) {
	TestSort(MergeSort, NewSortReverseSorted(), t)
}

func TestMergesortAllEqual(t *testing.T) {
	TestSort(MergeSort, NewSortAllEquals(), t)
}
