package random

import (
	"math"
	"time"
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
	return NewGeneratorWithSeed(time.Now().UTC().Nanosecond())
}

//Most commonly used generator to initialize with user defined seed
func NewGeneratorWithSeed(xSeed int) *Generator {
	return NewGenericRandomGenerator(22 , 1234 , xSeed )
}

//Generic random number generator where caller control all parameters
func NewGenericRandomGenerator(multiplier int, increment int, xSeed int) *Generator {
	return &Generator{multiplier: multiplier, increment: increment, xSeed: xSeed}
}

//Generate the random numbers array for given count
func (g *Generator) GetNumbers(count int)  []int {
	return g.GetNumbersWithinMax(math.MaxInt32, count)
}

//Generate the random numbers array for given count with maxvalue limit
func (g *Generator) GetNumbersWithinMax(maxValue int, count int)  []int {
	g.modulus = maxValue
	numbers := make([]int, 0)
	start := g.xSeed
	for i := 0; i < count; i++ {
		num := (g.multiplier*start+g.increment)%g.modulus
		numbers = append(numbers, num)
		start = num
	}
	return numbers
}
