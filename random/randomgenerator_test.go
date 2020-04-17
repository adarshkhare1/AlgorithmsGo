package random

import (
	"fmt"
	"testing"
)

func TestNewRandomGenerator(t *testing.T){
	/* read the list until head is not nil */
	g := NewRandomGenerator(1000,12,12,22)
	count := 10
	nums := g.GenerateNumbers(count)
	for i := 0; i < count; i++{
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
}
