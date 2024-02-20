package v5

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Sleeper - custom interface type
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper - we can adjust the sleep time
type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration) // sleep is a variable that holds a reference to a function
}

func (s *ConfigurableSleeper) Sleep() {
	// call the function with the param value
	s.SleepFunc(s.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
