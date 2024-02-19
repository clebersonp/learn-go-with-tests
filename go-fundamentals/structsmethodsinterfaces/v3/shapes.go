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

func Area(rec Rectangle) float64 {
	return rec.Width * rec.Height
}

func Area(circle Circle) float64 {
	return circle * math.Pi
}
