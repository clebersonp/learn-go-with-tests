package v3

import (
	"bytes"
	"testing"
)

// SpySleeper - custom type
// Spies are a kind of mock which can record how a dependency is used.
// They can record the arguments sent in, how many times it has been called, etc.
type SpySleeper struct {
	Calls int
}

// Sleep is an implementation of Sleeper (custom) interface
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	wantTimes := 3
	if spySleeper.Calls != wantTimes {
		t.Errorf("not enough calls to sleeper, want %d got %d", wantTimes, spySleeper.Calls)
	}
}
