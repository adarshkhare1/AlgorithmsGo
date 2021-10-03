package datastructures

type MaxHeap struct {
	items []int
}

func NewMaxHeap(items []int) *MaxHeap {
	return &MaxHeap{items: items}
}

func (h *MaxHeap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (h *MaxHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) isLeaf(index int) bool {
	size := len(h.items)
	if index >= size/2 && index <= size {
		return true
	}
	return false
}

func (h *MaxHeap) BuildMaxHeap() {
	size := len(h.items)
	for index := size/2-1; index >= 0; index-- {
		h.MaxHeapify(index, size)
	}
}

func (h *MaxHeap) MaxHeapify(index int, size int)  {
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
			h.MaxHeapify(largest, size)
		}
	}
}