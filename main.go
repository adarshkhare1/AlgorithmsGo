package Algorithms

import "Algorithms/random"

func main() {
	g := random.NewRandomGenerator()
	g.GetNumbers(10)
}