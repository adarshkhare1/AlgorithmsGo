package combinatorial

import (
	"fmt"
	"testing"
)

func TestNQueentProblem(t *testing.T) {
	maxBoardSize := 10
	for boardSize := 1; boardSize <= maxBoardSize; boardSize++ {
		prob := NewNQueenProblem(boardSize)
		combinations := prob.FindQueenPositions()
		if len(combinations) > 0 {
			for i := 0; i < len(combinations); i++ {
				fmt.Printf("\t %v\n", combinations[i])
			}
		} else {
			fmt.Printf("\t<<NONE>>\n")
		}
	}
}
