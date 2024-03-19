package v2

import (
	"math"
	"testing"
	"time"
)

// run a specific file to test
// go test -run "^TestSecondsIn*" -v

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	if want != got {
		t.Errorf("wanted %v radians, but got %v", want, got)
	}
}
