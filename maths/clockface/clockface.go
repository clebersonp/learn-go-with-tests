package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

// secondsInRadians One second is (2Pi / 60) radians
func secondsInRadians(t time.Time) float64 {
	// dividing by +Inf the result will be zero
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

// minutesInRadians One minute is (2Pi / 60) radians
func minutesInRadians(t time.Time) float64 {
	// dividing by +Inf the result will be zero
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / (float64(t.Minute()))))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

// hoursInRadians
func hoursInRadians(t time.Time) float64 {
	// dividing by +Inf the result will be zero
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / (float64(t.Hour() % hoursInClock))))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{X: x, Y: y}
}
