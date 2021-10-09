package searching

import (
	"math/rand"
	"time"
)

// FindIndexForItem Returns the index of value k in a sorted searchList, return -1 if k not found.
func FindIndexForItem(searchList []int, k int) int {
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
