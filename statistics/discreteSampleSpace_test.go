package statistics

import (
	. "Algorithms/random"
	"fmt"
	"testing"
)

func TestEvaluateChiSquareVariation(t *testing.T) {
	degreeOfFreedom := 6
	x := 0.0
	numIterations := 100 // Number of iterations in experiment
	nExperiments := 10 //number of times experiment is repeated
	for i := 0; i < nExperiments; i++ {
		g := NewRandomGenerator()
		experiment := NewDiscreteSampleSpace(degreeOfFreedom)
		for j := 0; j < numIterations; j++ {
			num, _ := g.NextNumberWithMax(degreeOfFreedom-1)
			experiment.AddSample(string(num))
		}
		x += experiment.EvaluateChiSquareVariation()
	}
	aveChiSquare := x/ float64(nExperiments)
	fmt.Printf("Average ChiSquare Variation for %d degree of freedom is %f.\n",
		degreeOfFreedom,
		aveChiSquare)
}
