package v2

import (
	"sync"
	"testing"
)

// run in terminal: go vet
// 'go vet' shows up errors because we are copying Mutex value after first use.
// doc: A Mutex must not be copied after first use.
// When we pass our Counter (by value) to assertCounter it will try and create a copy of the mutex
// To solve this we should pass in a pointer to our Counter instead, so change the signature of assertCounter
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("it runs concurrently", func(t *testing.T) {
		wantedCount := 1_000
		counter := NewCounter()

		/*
			A WaitGroup waits for a collection of goroutines to finish.
			The main goroutine calls Add to set the number of goroutines to wait for.
			Then each of the goroutines runs and calls Done when finished.
			At the same time, Wait can be used to block until all goroutines have finished.
		*/
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		// By waiting for wg.Wait() to finish before making our assertions
		// we can be sure all of our goroutines have attempted to Inc the Counter
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.value != want {
		t.Errorf("got %d, want %d", got.value, want)
	}
}
