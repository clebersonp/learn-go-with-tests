package main

import (
	"fmt"
	"runtime"
	"sync"
)

// https://en.wikipedia.org/wiki/Race_condition

func main() {

	// run to analyse race condition: go run -race .
	fmt.Println("============Before=============")
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	// the race condition problem occur even we use counter pointer(memory address)
	// the problem is that we are sharing a variable and manipulating its value concurrently
	counter := 0
	totalOfGoRoutines := 1_000

	// a way to sync all GoRoutines
	var wg sync.WaitGroup
	wg.Add(totalOfGoRoutines)

	for i := 0; i < totalOfGoRoutines; i++ {
		// create a go routine with keyword 'go' in front a function call
		go func() {
			// take the shared variable value and store it here inside anonymous function
			internal := counter

			// yield (communicate OS to take other process for while)
			runtime.Gosched()

			// increment internal variable by 1
			internal++

			// store internal value to shared variable counter
			counter = internal

			// communicate wait group that it's done
			wg.Done()
		}()
	}

	// communicate the main process (GoRoutine) that it need to wait for all go routines done
	wg.Wait()

	fmt.Println("============After=============")
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
