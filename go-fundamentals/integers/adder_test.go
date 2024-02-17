package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

// ExampleAdd is an example for Godoc
func ExampleAdd() {
	x, y := 2, 2

	fmt.Printf("x: %d, y: %d\n", x, y)
	fmt.Printf("sum: %d\n", Add(x, y))

	// https://go.dev/blog/examples
	// godoc -http=localhost:6060

	// Output:
	// x: 2, y: 2
	// sum: 4
}
