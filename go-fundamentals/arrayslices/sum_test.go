package arrayslices

import "testing"

func TestSum(t *testing.T) {
	// arrays have a fixed capacity
	// can be initialized in two ways:
	// 1 - [size]int{number, number, ...}
	// 2 - [...]int{number, number, ...}
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}
