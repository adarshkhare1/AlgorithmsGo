package Algorithms

import "Algorithms/random"

func main() {
	g := random.NewRandomGenerator()
	g.NextNumbersWithinMax(10, 10)
}