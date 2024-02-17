package iteration

import (
	"testing"
)

// To run tests and benchmarks = go test -v -bench=.

// BenchmarkRepeat Benchmarks in Go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

// BenchmarkParallelRepeat Benchmarks in Go
func BenchmarkParallelRepeat(b *testing.B) {
	parallel := func(pb *testing.PB) {
		for pb.Next() {
			Repeat("a", 0)
		}
	}
	b.RunParallel(parallel)
}

func TestRepeat(t *testing.T) {
	t.Run("zero times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "a"
		assertionHelper(t, repeated, expected)
	})
	t.Run("one time", func(t *testing.T) {
		repeated := Repeat("flow", 1)
		expected := "flow"
		assertionHelper(t, repeated, expected)
	})
	t.Run("two times", func(t *testing.T) {
		repeated := Repeat("flow", 2)
		expected := "flowflow"
		assertionHelper(t, repeated, expected)
	})
	t.Run("7 times", func(t *testing.T) {
		repeated := Repeat("ba", 7)
		expected := "bababababababa"
		assertionHelper(t, repeated, expected)
	})
}

func assertionHelper(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
