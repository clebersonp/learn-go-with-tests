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

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area - func (receiver type) name(...parameters type) return type {}
// This is a method of circle instance type
// func (receiverName ReceiverType) MethodName(args)
// When your method is called on a variable of that type,
// you get your reference to its data via the receiverName variable.
// In many other programming languages this is done implicitly, and you access the receiver via 'this'.
// It is a convention in Go to have the receiver variable be the first letter of the type
func (c Circle) Area() float64 {
	// Area = π × r^2
	return math.Pi * math.Pow(c.Radius, 2)
}
