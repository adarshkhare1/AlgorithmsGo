package random

import (
	"fmt"
	"testing"
	"time"
)

func TestGenericRandomGenerator(t *testing.T){
	/* read the list until head is not nil */
	g := NewRandomGenerator()
	nums := g.GetNumbers(32)
//	printGeneratorParameters(g)
	printNumberList(nums)
	fmt.Println()
}

func TestRandomGeneratorWithSeed(t *testing.T){
	//Using System tick count as seed
	g := NewGeneratorWithSeed(time.Now().UTC().Nanosecond())
	nums := g.GetNumbersWithinMax(100, 20)
	printNumberList(nums)
	fmt.Println()
}

func printNumberList(nums []int) {
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
}

func printGeneratorParameters(g *Generator) {
	fmt.Printf("seed      = %d\n", g.xSeed)
	fmt.Printf("multipler = %d\n",g.multiplier)
	fmt.Printf("modulus   = %d\n",g.modulus)
	fmt.Printf("increment = %d\n",g.increment)
}