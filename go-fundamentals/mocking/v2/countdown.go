package v2

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}
