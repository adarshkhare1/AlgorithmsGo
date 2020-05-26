package statistics

type ExperimentData struct {
	Results map[string]ExperimentResult
}

func NewExperimentData() *ExperimentData {
	return &ExperimentData{ Results: make(map[string]ExperimentResult) }
}

type ExperimentResult struct {
	Count int64
	ExpectedP float64
}

func NewExperimentResult() *ExperimentResult {
	return &ExperimentResult{}
}

func (data *ExperimentData) AddResult(id string) {
	r := data.Results[id]
	r.Count++
	data.Results[id] = r
}

//Initializes expected probability as equal for given degree of freedom, index will be the ID of the value.
func (data *ExperimentData)InitEqualDistributionResultsMap(degreeOfFreedom int)  {
	p := 1.00/float64(degreeOfFreedom)
	for i := 0; i < degreeOfFreedom; i++ {
		result := NewExperimentResult()
		result.ExpectedP = p
		result.Count = 0
		data.Results[string(i)] = *result
	}
}

//Returns the calculated probability distribution of experiment data
func (data *ExperimentData) FindProbabilityDistribution() map[string]float64 {
	sum := data.calculateTotalIterations()
	result := make(map[string]float64)
	for key, value := range data.Results {
		result[key] += float64(value.Count)/float64(sum)
	}
	return result
}

//Calculate total number of iterations in given experiment data
func (data *ExperimentData)calculateTotalIterations() int64  {
	var sum int64
	for _, value := range data.Results {
		sum += int64(value.Count)
	}
	return sum
}