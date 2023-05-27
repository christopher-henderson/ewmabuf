package ewma

import (
	"testing"
)

func TestStdDeviation_Compute(t *testing.T) {
	std := &StdDeviation{mean: &Average{}}
	std.Compute(4)
	std.Compute(7)
	std.Compute(1)
	std.Compute(3)
	std.Compute(6)
	std.Compute(8)
	std.Compute(2)
	std.Compute(9)
	got := std.Compute(5)
	want := 2.581988897471611
	if got != want {
		t.Fatalf("expected %.20f, got %.20f", want, got)
	}
}

func TestStdDeviation_ComputeEWMA(t *testing.T) {
	std := NewStdDeviation(NewEWMA(1, 0))
	std.Compute(4)
	std.Compute(7)
	std.Compute(1)
	std.Compute(3)
	std.Compute(6)
	std.Compute(8)
	std.Compute(2)
	std.Compute(9)
	got := std.Compute(5)
	want := 2.581988897471611
	if got != want {
		t.Fatalf("expected %.20f, got %.20f", want, got)
	}
}

func TestStdDeviation_ComputeEWMA2(t *testing.T) {
	std := NewStdDeviation2(0.5)
	deviation, mean := std.Compute(4)
	t.Logf("4 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(7)
	t.Logf("7 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(1)
	t.Logf("1 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(3)
	t.Logf("3 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(6)
	t.Logf("6 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(8)
	t.Logf("8 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(2)
	t.Logf("2 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(9)
	t.Logf("9 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)

	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)

	deviation, mean = std.Compute(100)
	t.Logf("100 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
	deviation, mean = std.Compute(99)
	t.Logf("99 => deviation: %f, mean: %f, shrink threshold: %f", deviation, mean, deviation+mean)
}

func TestStdDeviation_ComputeEWMA2WithBuf(t *testing.T) {
	//kib := 1024
	//bufSize := 64 * kib
	std := NewStdDeviation2(0.5)
	deviation, mean := std.Compute(4)
	t.Logf("4 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(7)
	t.Logf("7 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(1)
	t.Logf("1 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(3)
	t.Logf("3 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(6)
	t.Logf("6 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(8)
	t.Logf("8 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(2)
	t.Logf("2 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(9)
	t.Logf("9 => deviation: %f, mean: %f", deviation, mean)
	deviation, mean = std.Compute(5)
	t.Logf("5 => deviation: %f, mean: %f", deviation, mean)
}
