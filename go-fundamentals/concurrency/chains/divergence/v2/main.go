package main

import (
	"fmt"
	"sync"
	"time"
)

// to execute: go run .
// verify 'race condition': go run -race .

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	numbers := generate(30)

	go sender(numbers, channel1)

	// maximum of goroutines at a time
	throttling := 5
	go diverge(throttling, channel1, channel2)

	for i := range channel2 {
		fmt.Println(i)
	}
	//                       |			divergence			| convergence
	// chan1 send values to -> many goroutines send values to -> chan2
}

func sender(numbers []int, channel chan int) {
	defer close(channel)
	for n := range numbers {
		channel <- n
	}
}

func diverge(throttling int, ch1, ch2 chan int) {
	defer close(ch2)
	var wg sync.WaitGroup
	defer wg.Wait()

	// 1 - It diverges from ch1 values to many goroutines
	// 2 - After that, any goroutine send it value to ch2. Here we have a convergence, where all the scattered goroutines will send their values to a single channel
	// 3 - There is a limit to how many goroutine can be run at once, limited by the 'throttling' value
	for range throttling {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range ch1 {
				ch2 <- work(i)
			}
		}()
	}
}

func work(n int) int {
	time.Sleep(1e3 * time.Millisecond)
	return n
}

func generate(amount int) []int {
	var numbers []int
	for n := range amount {
		numbers = append(numbers, n)
	}
	return numbers
}
