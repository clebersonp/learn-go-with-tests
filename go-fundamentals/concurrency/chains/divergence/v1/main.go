package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// to execute: go run .
// verify 'race condition': go run -race .

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	numbers := generate(20)

	go sender(numbers, channel1)
	go diverge(channel1, channel2)

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

func diverge(ch1, ch2 chan int) {
	defer close(ch2)
	var wg sync.WaitGroup
	defer wg.Wait()

	// 1 - It diverges from ch1 values to many goroutines
	// 2 - After that, any goroutine send it value to ch2. Here we have a convergence, where all the scattered goroutines will send their values to a single channel
	for i := range ch1 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch2 <- work(n)
		}(i)
	}
}

func work(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2e3)))
	return n
}

func generate(amount int) []int {
	var numbers []int
	for n := range amount {
		numbers = append(numbers, n)
	}
	return numbers
}
