package statistics

import (
	. "Algorithms/random"
	"fmt"
	"testing"
)

func TestEvaluateChiSquareVariation(t *testing.T) {
	degreeOfFreedom := 6
	x := 0.0
	numIterations := 500 // Number of iterations in experiment
	nExperiments := 10 //number of times experiment is repeated
	for i := 0; i < nExperiments; i++ {
		g := NewRandomGenerator()
		d := NewExperimentData()
		d.InitEqualDistributionResultsMap(degreeOfFreedom)
		for j := 0; j < numIterations; j++ {
			num, _ := g.NextNumberWithMax(degreeOfFreedom-1)
			index := num
			r := d.Results[string(index)]
			r.Count++
			d.Results[string(index)] = r
		}
		variation := EvaluateChiSquareVariation(*d)
		x += variation
	}
	fmt.Printf("Average ChiSquare Variation with % d degree of freedom is %f.\n",
		degreeOfFreedom,
		x/float64(nExperiments))
}
