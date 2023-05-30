package ewma

import (
	"encoding/json"
	"os"
	"testing"
)

type StatsRun struct {
	Messages         []int
	BufSize          []int
	AverageOverAlloc []float64

	sum int
	avg float64

	buf *EwmaBuffer
}

func NewStatsRun(t *testing.T, path string) *StatsRun {
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var messages []int
	err = json.NewDecoder(f).Decode(&messages)
	if err != nil {
		t.Fatal(err)
	}
	return &StatsRun{
		Messages:         messages,
		BufSize:          make([]int, len(messages)),
		AverageOverAlloc: make([]float64, len(messages)),
		sum:              0,
		avg:              0,
		buf:              NewEwmaBuffer(0.01, 0.5, 1024*64),
	}
}

func (s *StatsRun) Run() {
	for i, message := range s.Messages {
		s.buf.BufferFor(message)
		s.BufSize[i] = s.buf.length()
		s.sum += s.buf.length() - message
		s.AverageOverAlloc[i] = float64(s.sum) / float64(i+1)
	}
}

func (s *StatsRun) Save(t *testing.T, path string) {
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(s)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAsdasd(t *testing.T) {
	stats := NewStatsRun(t, "/home/chris/projects/ewmabuf/normal.json")
	stats.Run()
	stats.Save(t, "/home/chris/projects/ewmabuf/normal_result.json")
}
