package Algorithms

import "Algorithms/random"

func main() {
	g := random.NewGenericRandomGenerator(5,2,4)
	g.GetNumbers(10)
}