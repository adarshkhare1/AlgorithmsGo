package sorting

//Quick sort impplementation using element swap.
func QuickSort(items []int) []int {

	if len(items) <= 1 {
		return items
	}
	quickSortElements(items, 0, len(items))
	return items
}

func quickSortElements(items []int, start int, end int) {

	if start < end {
		pivotIndex := partition(items, start, end)
		quickSortElements(items, start, pivotIndex)
		quickSortElements(items, pivotIndex+1, end)
	}
}

func partition(items []int, start int, end int)  int {
	pivot := items[end-1]
	i := start-1
	for j:= start; j < end - 1; j++{
		if items[j] < pivot{
			i++
			swapArrayElements(items, i, j)
		}
	}
	swapArrayElements(items, i+1, end-1)
	return i+1
}


