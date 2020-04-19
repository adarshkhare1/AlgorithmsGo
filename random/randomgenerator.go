package random

import (
	. "math"
	. "time"
)

// generate uniform distribution random numbers using Linear Congruential Method
//Algorithm -
//NumNext = (Multiplier*Num + increment) mod M
// Num = Numnext
type Generator struct {
	modulus int //modulus
	multiplier int //multiplier
	increment int //increment
	xSeed int // starting value
}

//Most commonly used generator to initialize with time tick as seed
func NewRandomGenerator() *Generator {
	return NewGeneratorWithSeed(Now().UTC().Nanosecond())
}

//Most commonly used generator to initialize with user defined seed
func NewGeneratorWithSeed(xSeed int) *Generator {
	//Using the default value of multiplier and increment as defined in ANSI C
	return &Generator{multiplier: 1103515245, increment: 12345, xSeed: xSeed}
}

//Generate the random numbers array for given count, range will be selected based on
//desired population.
func (g *Generator) GetNumbers(numbersCount int)  []int {
	max := MaxInt32
	return g.GetNumbersWithinMax(max, numbersCount)
}

//Generate the random numbers array for given count with maxvalue limit
//Number sequence will repeat if maxValue less than desired population size
func (g *Generator) GetNumbersWithinMax(maxValue int, numbersCount int)  []int {
	if maxValue < MaxInt32 {
		g.modulus = maxValue + 1
	} else {
		g.modulus = MaxInt32
	}
	if numbersCount < MaxInt32/2 {
		g.multiplier = nextPowerOf2(numbersCount) + 1
		if g.multiplier == g.modulus {
			g.multiplier = g.multiplier + 1
		}
	}
	numbers := make([]int, 0)
	start := g.xSeed
	for i := 0; i < numbersCount; i++ {
		num := (g.multiplier*start+g.increment)%g.modulus
		numbers = append(numbers, num%numbersCount)
		start = num
	}
	return numbers
}

func nextPowerOf2(n int) int {
	if n < 2 {
		return 0
	}
	k := 1
	for n != 0 {
		n = n >> 1
		k = k << 1
	}
	return k
}

