package statistics

import (
	"errors"
	"math"
)

type ContinuousSampleSpace struct {
	Results []float64
	MinimumValue float64
	MaximumValue float64
}

func NewContinuousSampleSpace() *ContinuousSampleSpace {
	return NewContinuousSampleSpaceWithRange(0.0, math.MaxFloat64)
}

func NewContinuousSampleSpaceWithRange(minValue float64, maxValue float64) *ContinuousSampleSpace {
	return &ContinuousSampleSpace{ Results: make([]float64,0), MinimumValue: minValue, MaximumValue: maxValue}
}

func (experiment *ContinuousSampleSpace) AddResult(val float64) error{
	if val < experiment.MinimumValue || val > experiment.MaximumValue {
		return errors.New("value is out of range")
	}
	experiment.Results = append(experiment.Results, val )
	return nil
}


func (experiment *ContinuousSampleSpace) EvaluateKSDistribution() error{
	// Not yet implemented

	return errors.New("not yet implemented")
}
