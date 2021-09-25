package searching

import "math"

type MaxSubArrayTestInput struct {
	Input       []int
	outStart    int
	outEnd int
	outSum int
}

func MaxSubArrayBasic() *MaxSubArrayTestInput {
	return &MaxSubArrayTestInput{
		Input:    []int{13,-3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
		outStart: 7,
		outEnd:   10,
		outSum:   43}
}

func MaxSubArrayFullLength() *MaxSubArrayTestInput {
	return &MaxSubArrayTestInput{
		Input:    []int{1, 4, 2, 5, 3},
		outStart: 0,
		outEnd:   4,
		outSum:   15}
}

func MaxSubArraySingleElement() *MaxSubArrayTestInput {
	return &MaxSubArrayTestInput{
		Input:    []int{3},
		outStart: 0,
		outEnd:   1,
		outSum:   3}
}
func MaxSubArrayEmpty() *MaxSubArrayTestInput {
	return &MaxSubArrayTestInput{
		Input:    []int{},
		outStart: 0,
		outEnd:   0,
		outSum:   math.MinInt32}
}

func MaxSubArrayAllNegative() *MaxSubArrayTestInput {
	return &MaxSubArrayTestInput{
		Input:    []int{-3, -25, -3, -16, -23, -7, -5, -22, -4},
		outStart: 2,
		outEnd:   3,
		outSum:   -19}
}


