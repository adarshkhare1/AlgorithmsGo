package sorting

import (
	"Algorithms/testHelper"
	"testing"
)

type SortTestInput struct {
	Input []int
	Expected []int
}

func NewSortSequence() *SortTestInput {
	return &SortTestInput{
		Input: []int{1, 4, 2, 5, 3, 6},
		Expected: []int{1, 2, 3, 4, 5, 6}}
}

func NewSortBasic() *SortTestInput {
	return &SortTestInput{
		Input: []int{5, 13, 2, 25, 7, 17, 20, 8, 4},
		Expected: []int{2, 4, 5, 7, 8, 13, 17, 20, 25}}
}

func NewSortEmpty() *SortTestInput {
	return &SortTestInput{
		Input: []int{},
		Expected: []int{}}
}

func NewSortPreSorted() *SortTestInput {
	return &SortTestInput{
		Input: []int{1, 2, 3, 4, 5},
		Expected: []int{1, 2, 3, 4, 5}}
}

func NewSortReverseSorted() *SortTestInput {
	return &SortTestInput{
		Input: []int{5, 4, 3, 2, 1},
		Expected: []int{1, 2, 3, 4, 5}}
}

func NewSortAllEquals() *SortTestInput {
	return &SortTestInput{
		Input: []int{5, 5, 5, 5, 5},
		Expected: []int{5, 5, 5, 5, 5}}
}

func TestSort(fp func(items []int) []int, in *SortTestInput, t *testing.T) {
	result := fp(in.Input)
	testHelper.VerifyArraysAreEqual(t, in.Expected, result)
}
