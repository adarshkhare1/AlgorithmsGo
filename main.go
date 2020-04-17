package Algorithms

import "Algorithms/random"

func main() {
	g := random.NewRandomGenerator(5,2,4,22)
	g.GenerateNumbers(10)
}