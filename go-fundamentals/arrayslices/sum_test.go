package arrayslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// arrays have a fixed capacity
		// can be initialized in two ways:
		// 1 - [size]int{number, number, ...}
		// 2 - [...]int{number, number, ...}
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertion(t, got, want, numbers)
	})

}

func assertion(t testing.TB, got, want int, given [5]int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, given)
	}
}
