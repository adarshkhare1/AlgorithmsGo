package sorting

import "testing"

func TestHeapSort(t *testing.T) {
	TestSort(HeapSort, NewSortBasic(), t)
}

func TestHeapSortSequence(t *testing.T) {
	TestSort(HeapSort, NewSortSequence(), t)
}

func TestHeapSortEmpty(t *testing.T) {
	TestSort(HeapSort, NewSortEmpty(), t)
}

func TestHeapSortPreSorted(t *testing.T) {
	TestSort(HeapSort, NewSortPreSorted(), t)
}

func TestHeapSortReverseSorted(t *testing.T) {
	TestSort(HeapSort, NewSortReverseSorted(), t)
}

func TestHeapSorttAllEqual(t *testing.T) {
	TestSort(HeapSort, NewSortAllEquals(), t)
}

