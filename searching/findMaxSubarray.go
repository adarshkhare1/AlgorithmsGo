package searching

import "math"

type SubArray struct {
	Array []int // Parent array
	Start int // start of subarray
	End int // end of subarray
	Sum int // Sum of subarray elements

}

func NewSubArray(array []int, start int, end int) *SubArray {
	return &SubArray{Array: array, Start: start, End: end,  Sum: math.MinInt32}
}

// FindMaxSubArray : Return the Subarray from the given array, where sum of
// the subarray members is maximum. Minimum SumArray size is 2 elements
// If single element array is given as input, sum will be the element value itself.
func FindMaxSubArray(array []int) SubArray {
	if len(array) <= 1 {
		result := *NewSubArray(array,0, len(array))
		if len(array) == 1 {
			result.Sum = array[0]
		}
		return result
	} else {
		return findMaxSubArray(*NewSubArray(array, 0, len(array)))
	}
}

func findMaxSubArray(inArray SubArray) SubArray {
	if inArray.End == inArray.Start{//inArray.Sum = inArray.Array[inArray.Start]
		return inArray
	} else {
		mid := (inArray.Start+inArray.End) / 2
		leftSubArray := findMaxSubArray(*NewSubArray(inArray.Array,inArray.Start, mid))
		rightSubArray := findMaxSubArray(*NewSubArray(inArray.Array,mid+1, inArray.End))
		crossSubArray := findMaxCrossingSubArray(inArray, mid)
		if leftSubArray.Sum >= rightSubArray.Sum && leftSubArray.Sum >= crossSubArray.Sum{
			return leftSubArray
		}else if rightSubArray.Sum >= leftSubArray.Sum && rightSubArray.Sum >= crossSubArray.Sum {
			return rightSubArray
		} else {
			return crossSubArray
		}
	}
}

func findMaxCrossingSubArray(array SubArray, mid int) SubArray {
	leftSum := math.MinInt32
	sum := 0
	maxStart := mid
	maxEnd  := mid
	for i:=mid; i >= array.Start; i-- {
		sum = sum + array.Array[i]
		if sum > leftSum {
			leftSum = sum
			maxStart = i
		}
	}
	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i < array.End; i++ {
		sum = sum + array.Array[i]
		if sum > rightSum {
			rightSum = sum
			maxEnd = i
		}
	}
	result := *NewSubArray(array.Array, maxStart, maxEnd)
	result.Sum = leftSum+rightSum
	return result
}