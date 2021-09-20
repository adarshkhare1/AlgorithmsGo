package sorting

func InsertionSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	j := 1
	for j < len(items) {
		key := items[j]
		// Insert items[j] into the sorted sequence
		i := j -1
		for i >= 0 && items[i] > key {
			items[i+1] = items[i]
			i--
		}
		items[i+1] = key
		j++
	}
	return items
}
