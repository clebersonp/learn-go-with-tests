package arrayslices

import (
	"reflect"
	"testing"
)

// Go's built-in testing toolkit features a coverage tool. In terminal type and enter: go test -v -cover

func assertion(t testing.TB, got, want int, given [5]int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, given)
	}
}

func assertionSlice(t testing.TB, got, want int, given []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, given)
	}
}

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

func TestSumSlice(t *testing.T) {
	t.Run("collection of 3 size with slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumSlice(numbers)
		want := 6
		assertionSlice(t, want, got, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Go does not let you use equality operators (!=) with slices.
	// You could write a function to iterate over each got and want slice and check their values
	// but for convenience’s sake, we can use reflect.DeepEqual function which is useful for seeing
	// if any two variables are the same
	// It's important to note that reflect.DeepEqual is not "type safe"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
