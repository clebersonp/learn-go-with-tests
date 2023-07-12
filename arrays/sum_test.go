package arrays

import (
	"reflect"
	"testing"
)

// To run cover: $ go test -cover

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// array can initialize as: [5]int{1, 2, 3, 4, 5}
		// array can initialize as: [...]int{1, 2, 3, 4, 5}
		// slice has no fixed size and can initialize as: []int{1, 2, 3, 4, 5}
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertResult(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("varargs collection of numbers to sum all", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertResult(t testing.TB, got int, want int, numbers []int) {
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
