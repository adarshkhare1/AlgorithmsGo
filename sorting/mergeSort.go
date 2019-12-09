package sorting

//TODO: Implement inplace sorting instead of creating new array
func merge(left []int, right []int) []int {

	result := make([]int, len(left) + len(right))
	i := 0
	j :=0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	for i < len(left) { result[i+j] = left[i]; i++ }
	for j < len(right) { result[i+j] = right[j]; j++ }
	return result

}

func Mergesort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	left := Mergesort(items[:middle])
	right := Mergesort(items[middle:])
	return merge(left, right)
}


