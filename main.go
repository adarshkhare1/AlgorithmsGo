package AlgorithmsGo

import "AlgorithmsGo/random"

func main() {
	g := random.NewRandomGenerator()
	g.NextNumbersWithinMax(10, 10)
}
