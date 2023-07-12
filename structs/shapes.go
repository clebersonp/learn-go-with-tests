package structs

import "math"

// struct does not explicitly implement interface methods, in Go is implicity
// The struct only needs to satisfy interface methods
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
