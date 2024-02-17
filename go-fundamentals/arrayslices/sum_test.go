package arrayslices

import (
	"slices"
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
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllWithAppend(t *testing.T) {
	got := SumAllWithAppend([]int{4, 1, 3}, []int{5, 8})
	want := []int{8, 13}

	// Go does not let you use equality operators (!=) with slices.
	// You could write a function to iterate over each got and want slice and check their values
	// but for convenience’s sake, we can use reflect.DeepEqual function which is useful for seeing
	// if any two variables are the same
	// It's important to note that reflect.DeepEqual is not "type safe"
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	// assigning a function to a variable
	// it's no different to assigning a variable to a string, or an int, functions in effect are values too
	checkSums := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum all tails with more than one element each", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("sum all tails with one element each", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{0})
		want := []int{0, 0}
		checkSums(t, got, want)
	})
	t.Run("sum all tails with at least one empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{7}, nil)
		want := []int{0, 0, 0}
		checkSums(t, got, want)
	})
	t.Run("sum all tails with sort number elements each slice", func(t *testing.T) {
		got := SumAllTails([]int{4, 7, 3}, []int{3}, []int{8, 1}, []int{9, 5, 3, 0}, []int{})
		want := []int{10, 0, 1, 8, 0}
		checkSums(t, got, want)
	})
}
