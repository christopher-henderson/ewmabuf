package ewma

import (
	"math"
)

type EwmaBuffer struct {
	//buf               []byte
	buf               int
	mean              float64
	variance          float64
	standardDeviation float64

	alpha        float64
	resizeFactor float64

	recentResize int
}

func NewEwmaBuffer(alpha float64, resizeFactor float64, initial int) *EwmaBuffer {
	return &EwmaBuffer{alpha: alpha, resizeFactor: resizeFactor, buf: initial, mean: float64(initial)}
}

func (e *EwmaBuffer) BufferFor(requested int) []byte {
	e.updateStatistics(float64(requested))
	newBufSize := e.length()
	if e.lengthf64() > e.mean+e.standardDeviation {
		// We are trending towards requiring smaller buffers
		// and need to resize downwards accordingly.
		newBufSize = int(math.Floor(e.mean + e.standardDeviation*e.resizeFactor))
	}
	if newBufSize < requested {
		newBufSize = int(math.Floor(float64(requested) + e.standardDeviation*e.resizeFactor))
	}
	if newBufSize != e.length() {
		e.resize(newBufSize)
	}
	return make([]byte, 0)
	//return e.buf[:requested]
}

func (e *EwmaBuffer) updateStatistics(length float64) {
	//e.mean = e.alpha*length + (1.0-e.alpha)*e.mean
	diff := length - e.mean
	incr := e.alpha * diff
	e.mean = e.mean + incr
	e.variance = (1 - e.alpha) * (e.variance + diff*incr)
	e.standardDeviation = math.Sqrt(e.variance)
}

func (e *EwmaBuffer) resize(target int) {
	if target < e.length() && e.recentResize > 0 {
		e.recentResize -= 1
		return
	}
	//e.buf = make([]byte, target)
	e.buf = target
	e.recentResize = 10
}

func (e *EwmaBuffer) length() int {
	//return len(e.buf)
	return e.buf
}

func (e *EwmaBuffer) lengthf64() float64 {
	//return float64(len(e.buf))
	return float64(e.buf)
}
