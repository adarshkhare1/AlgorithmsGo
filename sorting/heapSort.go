package sorting

func HeapSort(items []int) []int {
	buildMaxHeap(items)
	reverseArray(items)
	return items
}

func reverseArray(items []int) {
	start := 0
	for end := len(items) - 1; end >= len(items)/2; end-- {
		temp := items[end]
		items[end] = items[start]
		items[start] = temp
		start++
	}
}

func buildMaxHeap(items []int) {
	index := len(items)/2
	for index >= 0 {
		maxHeapify(items, index)
		index--
	}
}

func maxHeapify(items []int, index int)  {
	left := index+1
	right := index+2
	largest := index
	heapSize := len(items)
	if left < heapSize && items[left] > items[index] {
		largest = left
	} else {
		largest = index
	}
	if right < heapSize && items[right] > items[largest]{
		largest = right
	}
	if largest != index {
		temp := items[index]
		items[index] = items[largest]
		items[largest] = temp
		maxHeapify(items, largest)
	}
}