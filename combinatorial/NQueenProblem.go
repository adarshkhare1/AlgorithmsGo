package combinatorial

import "fmt"

// NQueenProblem is a classic problem of n queens. In how many ways can n Queens be placed
// on NxN board so that no two are in same row, column, or diagonal
type NQueenProblem struct {
	boardSize int
}

func NewNQueenProblem(boardSize int) *NQueenProblem {
	return &NQueenProblem{boardSize: boardSize}
}

// FindQueenPositions Return list of combinations where we can place queen in each row
// it will ignore the combinations where all rows can't be filled.
// return array is two-dimensional list, where each array of integers is list of columns for
// each row.  Row and column numbers are indexed 0.
// Example  [[1 3 0 2] [2 0 3 1]] are  two possible combinations,
//  In C=[1 3 0 2] C[i] is the column of i th row in this combination.
//  That means queens can be positioned (row, col) [(0,1), (1,3), (3,0), (4,2)
func (prob *NQueenProblem) FindQueenPositions() [][]int {
	fmt.Printf("NQueenProblem: Positions for board %dx%d >>\n", prob.boardSize, prob.boardSize)
	combinations := make([][]int, 0)
	for selectedRow := 0; selectedRow < prob.boardSize; selectedRow++ {
		for pos0 := 0; pos0 < prob.boardSize; pos0++ {
			positions := make([]int, prob.boardSize)
			positions[selectedRow] = pos0
			goodCombination := true
			for rowIndex := 0; rowIndex < prob.boardSize; rowIndex++ {
				if rowIndex == selectedRow {
					if !prob.validatePosition(positions, rowIndex, positions[rowIndex]) {
						goodCombination = false
						break
					}
				} else {
					colToAdd := positions[rowIndex]
					for colIndex := 0; colIndex < prob.boardSize; colIndex++ {
						colToAdd = colIndex
						if prob.validatePosition(positions, rowIndex, colToAdd) {
							positions[rowIndex] = colToAdd
							break
						} else if colIndex == prob.boardSize-1 {
							goodCombination = false
							break
						}
					}
				}
			}
			if goodCombination {
				combinations = append(combinations, positions)
			}
		}
	}
	return combinations
}

func (prob *NQueenProblem) validatePosition(positions []int, rowToValidate int, colToAdd int) bool {
	for rowIndex := 0; rowIndex < rowToValidate; rowIndex++ {
		colIndex := positions[rowIndex]
		if colToAdd == colIndex || prob._absInt(colToAdd-colIndex) == (rowToValidate-rowIndex) {
			return false
		}
	}
	return true
}

func (prob *NQueenProblem) _absInt(x int) int {
	return prob._absDiffInt(x, 0)
}

func (prob *NQueenProblem) _absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
