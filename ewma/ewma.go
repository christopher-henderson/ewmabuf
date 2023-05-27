package ewma

import "math"

type Mean interface {
	Mean(data float64) float64
}

type EWMA struct {
	alpha float64
	avg   float64
}

func NewEWMA(alpha float64, initial float64) *EWMA {
	return &EWMA{
		alpha: alpha,
		avg:   initial,
	}
}

func (e *EWMA) Mean(value float64) float64 {
	e.avg = e.alpha*value + (1.0-e.alpha)*e.avg
	return e.avg
}

type StdDeviation struct {
	mean       Mean
	population float64
	squareSum  float64
}

func NewStdDeviation(mean Mean) *StdDeviation {
	return &StdDeviation{mean: mean}
}

func (s *StdDeviation) Compute(data float64) float64 {
	mean := s.mean.Mean(data)
	s.population += 1
	s.squareSum += math.Pow(data, 2)
	variance := s.squareSum/s.population - math.Pow(mean, 2)
	return math.Sqrt(variance)
}

type Average struct {
	sum        float64
	population float64
}

func (a *Average) Mean(data float64) float64 {
	a.sum += data
	a.population += 1
	return a.sum / a.population
}

type StdDeviation2 struct {
	mean     *EWMA
	alpha    float64
	variance float64
}

func NewStdDeviation2(alpha float64) *StdDeviation2 {
	return &StdDeviation2{mean: NewEWMA(alpha, 0), alpha: alpha}
}

func (s *StdDeviation2) Compute(data float64) (float64, float64) {
	mean := s.mean.Mean(data)
	diff := data - mean
	incr := s.alpha * diff
	mean = mean + incr
	s.variance = (1 - s.alpha) * (s.variance + diff*incr)
	return math.Sqrt(s.variance), mean
}

func (s *StdDeviation2) Compute2(data float64) (float64, float64) {
	mean := s.mean.Mean(data)
	diff := data - mean
	incr := s.alpha * diff
	mean = mean + incr
	s.variance = (1 - s.alpha) * (s.variance + diff*incr)
	return math.Sqrt(s.variance), mean
}
