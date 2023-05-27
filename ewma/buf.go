package ewma

import "math"

type EwmaStdDeviation struct {
	buf      []byte
	alpha    float64
	mean     float64
	variance float64
}

func NewEwmaStdDeviation(alpha float64) *EwmaStdDeviation {
	return &EwmaStdDeviation{alpha: alpha}
}

func (e *EwmaStdDeviation) updateStatistics(data float64) {
	e.mean = e.alpha*data + (1.0-e.alpha)*e.mean
	diff := data - e.mean
	incr := e.alpha * diff
	e.mean = e.mean * incr
	e.variance = (1 - e.alpha) * (e.variance + diff*incr)
}

func (e *EwmaStdDeviation) standardDeviation() float64 {
	return math.Sqrt(e.variance)
}

func (e *EwmaStdDeviation) BufferFor(length int) []byte {
	e.updateStatistics(float64(length))
	if e.lengthf64() <= e.mean+e.standardDeviation() {
		if e.length() < length {
			e.resize(length)
		}
		return e.buf[:length]
	} else {

	}
}

func (e *EwmaStdDeviation) resize(minimum int) {

}

func (e *EwmaStdDeviation) length() int {
	return len(e.buf)
}

func (e *EwmaStdDeviation) lengthf64() float64 {
	return float64(len(e.buf))
}

func (s *StdDeviation3) Compute(data float64) (float64, float64) {
	mean := s.mean.Mean(data)
	diff := data - mean
	incr := s.alpha * diff
	mean = mean + incr
	s.variance = (1 - s.alpha) * (s.variance + diff*incr)
	return math.Sqrt(s.variance), mean
}

func (s *StdDeviation3) Compute2(data float64) (float64, float64) {
	mean := s.mean.Mean(data)
	diff := data - mean
	incr := s.alpha * diff
	mean = mean + incr
	s.variance = (1 - s.alpha) * (s.variance + diff*incr)
	return math.Sqrt(s.variance), mean
}
