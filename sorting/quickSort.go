package sorting

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
			swapElements(items, i, j)
		}
	}
	swapElements(items, i+1, end-1)
	return i+1
}

func swapElements(items []int, i int, j int) {
	if i < len(items) && j < len(items) && i != j {
		temp := items[i]
		items[i] = items[j]
		items[j] = temp
	}
}

