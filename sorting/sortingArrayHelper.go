package sorting

//Swap first and second index elements in items array
func swapArrayElements(items []int, first int, second int) {
	if first < len(items) && second < len(items) && first != second {
		temp := items[first]
		items[first] = items[second]
		items[second] = temp
	}
}