package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// iterate n times about the new received channel and get elements
	convergence := converge(hardWork("upload-images"), hardWork("convert-pdf"))

	for i := 0; i < 30; i++ {
		fmt.Println(<-convergence) // receiving the converged two channels values
	}
}

// hardWork simulate a hard work
func hardWork(name string) chan string {
	ch := make(chan string)
	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%v \t- %d", s, i)
			// take time to do the work
			time.Sleep(time.Duration(rand.Int64N(1e3)) * time.Millisecond)
		}
	}(name, ch)
	return ch
}

// convergence converge two channels into one
func converge(a, b chan string) chan string {
	// this is a closure function because it's passing other variables out of it scope
	convergence := make(chan string)
	go func() {
		for {
			convergence <- <-a
		}
	}()
	go func() {
		for {
			convergence <- <-b
		}
	}()
	return convergence
}
