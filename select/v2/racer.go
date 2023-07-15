package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string, err error) {
	// fmt.Printf("URL_A: %s\nURL_B: %s\n", urlA, urlB) // just to debug
	// synchronising processes
	// myVar := <-ch will blocking until receive a value.
	// select allows you to wait on multiple channels at once and the first one to send a value "wins"
	// and the code underneath the case is executed.
	select {
	// will be call func ping in the order that occur, but inside fun ping will run a new goroutine
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

// use a chan struct{} because we don't care about the value it will send
// we only care about the signal and it is the smallest data type available from a memory
// perspective so we get no allocation
func ping(url string) chan struct{} {
	// fmt.Printf("Sending request for %s\n", url) // just to debug
	// using make instead var to create channel is the correct, because with var will be 'zero value', and zero value for
	// channel is nil, and will blocking forever because you cannot send to nill channels
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
