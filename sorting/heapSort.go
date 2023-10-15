package sorting

import (
	. "AlgorithmsGo/datastructures"
)

func HeapSort(items []int) []int {
	heap := NewMaxHeap(items)
	heap.BuildMaxHeap()
	for i := len(items) - 1; i > 0; i-- {
		// Move current root to end
		swapArrayElements(items, 0, i)
		heap.MaxHeapify(0, i)
	}
	return items
}
