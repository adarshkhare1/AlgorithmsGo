package statistics

import "math"

func EvaluateChiSquareVariation(data ExperimentData) float64{
	sum := data.calculateTotalIterations()
	var result float64
	for _, value := range data.Results {
		expectedCount := int64(value.ExpectedP * float64(sum))
		result += math.Pow(float64(expectedCount-value.Count), 2)/float64(expectedCount)
	}
	return result/float64(sum)
}
