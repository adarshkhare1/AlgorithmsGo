package sorting

func merge(a []int, b []int) []int {

	result := make([]int, len(a) + len(b))
	i := 0
	j :=0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result[i+j] = a[i]
			i++
		} else {
			result[i+j] = b[j]
			j++
		}
	}

	for i < len(a) { result[i+j] = a[i]; i++ }
	for j < len(b) { result[i+j] = b[j]; j++ }
	return result

}

func Mergesort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	var a = Mergesort(items[:middle])
	var b = Mergesort(items[middle:])
	return merge(a, b)
}


