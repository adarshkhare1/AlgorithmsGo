package random

import (
	"fmt"
	"testing"
	"time"
)

func TestGenericRandomGenerator(t *testing.T){
	/* read the list until head is not nil */
	g := NewRandomGenerator()
	nums := g.GetNumbers(10)
	printNumberList(nums)
	fmt.Println()
}

func TestRandomGeneratorWithSeed(t *testing.T){
	//Using System tick count as seed
	g := NewGeneratorWithSeed(time.Now().UTC().Nanosecond())
	nums := g.GetNumbersWithinMax(1000, 10)
	printNumberList(nums)
	fmt.Println()
}

func printNumberList(nums []int) {
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
}