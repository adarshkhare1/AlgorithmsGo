package statistics

import (
	. "Algorithms/random"
	"fmt"
	"testing"
)

func TestEvaluateKSVariation(t *testing.T) {
	degreeOfFreedom := 10
	numIterations := 50 // Number of iterations in experiment
	x := 0.0
	nExperiments := 10 //number of times experiment is repeated
	for i := 0; i < nExperiments; i++ {
		g := NewRandomGenerator()
		experiment := NewContinuousSampleSpace()
		for j := 0; j < numIterations; j++ {
			num, _ := g.NextNumberWithMax(degreeOfFreedom-1)
			experiment.AddResult(float64(num))
		}
		experiment.EvaluateKSDistribution()
	}
	fmt.Printf("Average Kolmogorov-Smirnov Variation is %f.\n",
		x/float64(nExperiments))
}

