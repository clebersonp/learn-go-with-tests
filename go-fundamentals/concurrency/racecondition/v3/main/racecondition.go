package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// https://en.wikipedia.org/wiki/Race_condition

func main() {

	// run to analyse race condition: go run -race .
	fmt.Println("============Before=============")
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())

	// the race condition problem occur even we use counter pointer(memory address)
	// the problem is that we are sharing a variable and manipulating its value concurrently
	var counter int64
	totalOfGoRoutines := 1_000

	// a way to sync all GoRoutines
	var wg sync.WaitGroup
	wg.Add(totalOfGoRoutines)

	for i := 0; i < totalOfGoRoutines; i++ {
		// create a go routine with keyword 'go' in front a function call
		go func() {

			// another way to protect of 'race condition' is to use atomic package functions
			atomic.AddInt64(&counter, 1)
			// yield (communicate OS to take other process for while)
			runtime.Gosched()
			// get the counter value using atomic package
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
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
