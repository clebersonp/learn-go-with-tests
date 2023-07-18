package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		// WaitGroup is a means of waiting for goroutines to finish jobs
		// like join() of java
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

// Run : $ go vet
// to see warning like that: call of assertCounter copies lock value: v1.Counter contains sync.Mutex
// the doc says: A Mutex must not be copied after first use.
func assertCounter(t testing.TB, c *Counter, want int) {
	t.Helper()
	if c.Value() != want {
		t.Errorf("got %d, but want %d", c.Value(), want)
	}
}
