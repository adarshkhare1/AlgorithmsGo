package sorting

import (
	"testing"
)
func TestQuicksort(t *testing.T) {
	TestSort(QuickSort, NewSortBasic(), t)
}

func TestQuickSortSequence(t *testing.T) {
	TestSort(HeapSort, NewSortSequence(), t)
}

func TestQuickSortEmpty(t *testing.T) {
	TestSort(QuickSort, NewSortEmpty(), t)
}

func TestQuickSortPreSorted(t *testing.T) {
	TestSort(QuickSort, NewSortPreSorted(), t)
}

func TestQuickSortReverseSorted(t *testing.T) {
	TestSort(QuickSort, NewSortReverseSorted(), t)
}

func TestQuickSortAllEqual(t *testing.T) {
	TestSort(QuickSort, NewSortAllEquals(), t)
}
