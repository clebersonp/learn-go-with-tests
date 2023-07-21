package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	hourHandLength   = 50
	minuteHandLength = 80
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

// SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, "\n")
	io.WriteString(w, bezel)
	io.WriteString(w, "\n")
	hourHand(w, t)
	io.WriteString(w, "\n")
	minuteHand(w, t)
	io.WriteString(w, "\n")
	secondHand(w, t)
	io.WriteString(w, "\n")
	io.WriteString(w, svgEnd)
	io.WriteString(w, "\n")
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, line, p.X, p.Y, "#f00")
}

// minuterHand is the unit vector of the minute hand of an analogue clock at time `t`
// represented as a Point
func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, line, p.X, p.Y, "#000")
}

// hourHand is the unit vector of the hour hand of an analogue clock at time `t`
// represented as a Point
func hourHand(w io.Writer, t time.Time) {
	p := makeHand(hourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, line, p.X, p.Y, "#000")
}

func makeHand(p Point, length float64) Point {
	p = Point{X: p.X * length, Y: p.Y * length}                // scale
	p = Point{X: p.X, Y: -p.Y}                                 // flip
	return Point{X: p.X + clockCentreX, Y: p.Y + clockCentreY} // translate
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const line = `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:3px;"/>`

const svgEnd = `</svg>`
