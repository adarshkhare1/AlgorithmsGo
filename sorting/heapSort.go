package sorting

type maxHeap struct {
	items []int
}

func newMaxHeap(items []int) *maxHeap {
	return &maxHeap{items: items}
}

func (h *maxHeap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (h *maxHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (h *maxHeap) isLeaf(index int) bool {
	size := len(h.items)
	if index >= size/2 && index <= size {
		return true
	}
	return false
}

func (h *maxHeap) buildMaxHeap() {
	size := len(h.items)
	for index := size/2-1; index >= 0; index-- {
		h.maxHeapify(index, size)
	}
}

func (h *maxHeap) maxHeapify(index int, size int)  {
	if !h.isLeaf(index) {
		left := h.leftChildIndex(index)
		right := h.rightChildIndex(index)
		largest := index
		if left < size && h.items[left] > h.items[index] {
			largest = left
		}
		if right < size && h.items[right] > h.items[largest] {
			largest = right
		}
		if largest != index {
			swapArrayElements(h.items, index, largest)
			h.maxHeapify(largest, size)
		}
	}
}

func HeapSort(items []int) []int {
	heap := newMaxHeap(items)
	heap.buildMaxHeap()
	for i := len(items) - 1; i > 0; i-- {
		// Move current root to end
		swapArrayElements(items, 0, i)
		heap.maxHeapify(0, i)
	}
	return items
}
