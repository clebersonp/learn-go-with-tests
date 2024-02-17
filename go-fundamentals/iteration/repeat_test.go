package iteration

import "testing"

// To run tests and benchmarks = go test -v -bench=.

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// BenchmarkRepeat Benchmarks in Go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

// BenchmarkParallelRepeat Benchmarks in Go
func BenchmarkParallelRepeat(b *testing.B) {
	parallel := func(pb *testing.PB) {
		for pb.Next() {
			Repeat("a")
		}
	}
	b.RunParallel(parallel)
}
