package v3

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

// Area - func (receiver type) name(...parameters type) return type {}
// This is a method of circle instance type
func (circle Circle) Area() float64 {
	// Area = π × r^2
	return math.Pi * math.Pow(circle.Radius, 2)
}
