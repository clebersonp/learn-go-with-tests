package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("with seven repetition", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"
		assertResult(t, repeated, expected)
	})
	t.Run("with zero repetition", func(t *testing.T) {
		repeated := Repeat("b", 0)
		expected := ""
		assertResult(t, repeated, expected)
	})
	t.Run("with negative repetition", func(t *testing.T) {
		repeated := Repeat("c", -1)
		expected := "c"
		assertResult(t, repeated, expected)
	})
}

// prefix with Benchmark to write benchmark tests
// To run: $ go test -bench=.

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertResult(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

func ExampleRepeat() {
	fmt.Println("ba" + Repeat("na", 2))
	// Output: banana
}