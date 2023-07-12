package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, s Shape, hasArea float64) {
		t.Helper()
		got := s.Area()
		if got != hasArea {
			t.Errorf("%#v got %g, want %g", s, got, hasArea)
		}
	}

	// Table driven tests with anonymous struct
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// Run specific table test: $ go test -run TestArea/Rectangle -v
		{name: "Rectangle", shape: Rectangle{4.0, 2.0}, hasArea: 8.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.hasArea)
		})
	}
}
