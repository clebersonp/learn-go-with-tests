package clockface

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

// https://xml-to-go.github.io/
type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

// to run a specifc file needs to pass all the files involved
// $ go test clockface_assembly_test.go svgWriter.go clockface.go

// or by the pattern test name like:
// $ go test -run TestSVGWriter -v

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(0, 0, 0),
			line: Line{X1: 150, Y1: 150, X2: 150, Y2: 60},
		},
		{
			time: simpleTime(0, 0, 30),
			line: Line{X1: 150, Y1: 150, X2: 150, Y2: 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Lines) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines output %+v", line, svg.Lines)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(0, 0, 0),
			line: Line{X1: 150, Y1: 150, X2: 150, Y2: 70},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Lines) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Lines)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			time: simpleTime(6, 0, 0),
			line: Line{X1: 150, Y1: 150, X2: 150, Y2: 200},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Lines) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Lines)
			}
		})
	}
}

func containsLine(want Line, lines []Line) bool {
	for _, line := range lines {
		if want == line {
			return true
		}
	}
	return false
}
