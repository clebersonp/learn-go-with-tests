package v4

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Sleeper - custom interface type
type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
