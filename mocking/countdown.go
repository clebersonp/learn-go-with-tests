package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

// satisfy Sleeper interface
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(d time.Duration)
}

// satisfy Sleeper interface
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const (
	finalWorld     = "Go!"
	countdownStart = 3
)

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWorld)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
