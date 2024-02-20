package v5

import (
	"bytes"
	"reflect"
	"testing"
)

// SpyCountdownOperations - new custom spy to check all operations and their order
type SpyCountdownOperations struct {
	Calls []string
}

// Sleep implemented Sleeper interface method
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write implemented io.Writer interface method
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// constants to put them in the Calls field of SpyCountdownOperations
const (
	sleep = "sleep"
	write = "write"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}

		// as spySleeperPrinter implements both Write method from io.Writer and
		// Sleep() method from SpyCountdownOperations, we can pass both in Countdown function
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}

	})
}
