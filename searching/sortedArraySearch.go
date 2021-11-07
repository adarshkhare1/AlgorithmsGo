package searching

import (
	"math/rand"
	"time"
)

// FindItemIndexSequentialSkipSearch Returns the index of value k in a sorted searchList, return -1 if k not found.
func FindItemIndexSequentialSkipSearch(searchList []int, k int) int {
	i := 0
	size := len(searchList)
	rand.Seed(time.Now().UnixNano())
	for i < size && searchList[i] < k {
		j := i+rand.Intn(size-i)
		if searchList[i] < searchList[j] && searchList[j] <= k {
			i = j
			if searchList[i] == k {
				return i
			}
		}
		i++
	}
	if i >= size || searchList[i] > k {
		return -1
	}else {
		return i
	}
}

// FindItemIndexBinarySearch Returns the index of value k in a sorted searchList, return -1 if k not found.
//Method uses binary search.
func FindItemIndexBinarySearch(searchList []int, key int) int {
	return findItemViaBinarySearch(searchList, 0, len(searchList)-1, key)
}

// function recursively applies binary search on a sorted list.
func findItemViaBinarySearch(searchList []int, start, end int, key int) int {
	if (end - start) <=1 {
		if searchList[start] == key{
			return start
		} else if searchList[end] == key{
			return end
		}
	}else {
		mid := start + (end-start)/2
		if searchList[mid] == key {
			return mid
		} else if key < searchList[mid] {
			return findItemViaBinarySearch(searchList, start, mid, key)
		} else {
			return findItemViaBinarySearch(searchList, mid, end, key)
		}
	}
	return -1
}