package random

type Generator struct {
	modulus int //modulus
	multiplier int //multiplier
	increment int //increment
	xSeed int // starting value
}

func NewRandomGenerator(modulus int, multiplier int, increment int, xSeed int) *Generator {
	return &Generator{modulus: modulus, multiplier: multiplier, increment: increment, xSeed: xSeed}
}

// generate uniform distribution random numbers using Linear Congruential Method
//Algorithm -
//NumNext = (Multiplier*Num + increment) mod M
// Num = Numnext
func (g *Generator) GenerateNumbers(count int)  []int {
	numbers := make([]int, 0)
	start := g.xSeed
	for i := 0; i < count; i++ {
		num := (g.multiplier*start+g.increment)%g.modulus
		numbers = append(numbers, num)
		start = num
	}
	return numbers
}
