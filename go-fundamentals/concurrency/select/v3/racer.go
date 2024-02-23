package racer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	// we don't care what type is sent to the channel,
	// we just want to signal we are done and closing the channel works perfectly!
	// a chan struct{} is the smallest data type available from a memory perspective,
	// so we get no allocation versus a bool.
	// since we are closing and not sending anything on the chan, whi allocate anything?
	// For channels the zero value is nil and if you try and send to it with <- it will block forever because you
	// cannot send to nil channels. Always 'make' channels instead of 'var'
	ch := make(chan struct{})
	// we need a goroutine to pass value to channel
	go func(u string) {
		http.Get(u)
		close(ch) // close the channel because we don't send more values to it and to avoid another goroutine send signal
	}(url)
	return ch
}
