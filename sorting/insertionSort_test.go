package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	TestSort(InsertionSort, NewSortBasic(), t)
}

func TestInsertionSortEmpty(t *testing.T) {
	TestSort(InsertionSort, NewSortEmpty(), t)
}

func TestInsertionSortPreSorted(t *testing.T) {
	TestSort(InsertionSort, NewSortPreSorted(), t)
}

func TestInsertionSortReverseSorted(t *testing.T) {
	TestSort(InsertionSort, NewSortReverseSorted(), t)
}

func TestInsertionSortAllEqual(t *testing.T) {
	TestSort(InsertionSort, NewSortAllEquals(), t)
}
