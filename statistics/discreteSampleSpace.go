package statistics

import "math"

type DiscreteSampleSpace struct {
	ResultCount map[string]int64
	ExpectedP   map[string]float64
}

func NewDiscreteSampleSpace(degreeOfFreedom int) *DiscreteSampleSpace {
	experiment :=  &DiscreteSampleSpace{ ResultCount: make(map[string]int64), ExpectedP: make(map[string]float64)}
	experiment.initUniformDistributionMap(degreeOfFreedom)
	return experiment
}

//Initializes expected probability as equal for given degree of freedom, index will be the ID of the value.
func (experiment *DiscreteSampleSpace) initUniformDistributionMap(degreeOfFreedom int)  {
	p := 1.00/float64(degreeOfFreedom)
	for i := 0; i < degreeOfFreedom; i++ {
		experiment.ResultCount[string(i)] = 0
		experiment.ExpectedP[string(i)] = p
	}
}


func (experiment *DiscreteSampleSpace) AddSample(sampleId string) {
	experiment.ResultCount[sampleId]++
}

func (experiment *DiscreteSampleSpace) EvaluateChiSquareVariation() float64{
	totalIterations := experiment.getTotalIterations()
	var result float64
	for key, value := range experiment.ResultCount {
		result += calculateChiSquareFunction(totalIterations, experiment.ExpectedP[key], value)
	}
	return result/float64(totalIterations)
}

//Returns the calculated probability distribution of experiment data
func (experiment *DiscreteSampleSpace) CalculateProbabilityDistribution() map[string]float64 {
	sum := experiment.getTotalIterations()
	result := make(map[string]float64)
	for key, value := range experiment.ResultCount {
		result[key] += float64(value)/float64(sum)
	}
	return result
}

//Calculate total number of iterations in given experiment data
func (experiment *DiscreteSampleSpace) getTotalIterations() int64  {
	var sum int64
	for _, value := range experiment.ResultCount {
		sum += int64(value)
	}
	return sum
}

func calculateChiSquareFunction(totalIterations int64, expectedP float64, sampleCount int64) float64 {
	expectedCount := int64(expectedP * float64(totalIterations))
	return math.Pow(float64(expectedCount-sampleCount), 2)/float64(expectedCount)
}