package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected %q but got %q", expected, sum)
	}
}
