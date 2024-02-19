package v5

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// table driven tests.
	// Are useful when you want to build a list of test cases that can be tested in the same manner
	areaTableTest := []struct {
		shape Shape
		want  float64
	}{
		{
			shape: Rectangle{
				Width:  12.0,
				Height: 6.0,
			},
			want: 72.0,
		},
		{
			shape: Circle{Radius: 10.0},
			want:  314.1592653589793,
		},
		{
			shape: Triangle{12.0, 6.0},
			want:  36.0,
		},
	}

	for _, table := range areaTableTest {
		got := table.shape.Area()
		if got != table.want {
			t.Errorf("got %.2f want %.2f", got, table.want)
		}
	}
}
