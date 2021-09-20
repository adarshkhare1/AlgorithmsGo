package sorting

type SortTestInput struct {
	Input []int
	Expected []int
}

func NewSortBasic() *SortTestInput {
	return &SortTestInput{
		Input: []int{1, 4, 2, 5, 3},
		Expected: []int{1, 2, 3, 4, 5}}
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
