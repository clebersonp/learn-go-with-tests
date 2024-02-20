package main

import (
	mockingv5 "example.com/go-with-tests/go-fundamentals/mocking/v5"
	"os"
	"time"
)

func main() {
	sleeper := &mockingv5.ConfigurableSleeper{Duration: 4 * time.Second, SleepFunc: time.Sleep}
	mockingv5.Countdown(os.Stdout, sleeper)
}
