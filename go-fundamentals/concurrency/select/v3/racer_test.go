package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	// start a test server to simulate a response with delay
	slowServer := makeDelayedServer(5 * time.Millisecond)
	// start a test server to simulate the faster response without delay
	fastServer := makeDelayedServer(0 * time.Millisecond)

	// defer to close servers at the end of performing this block of code (TestRacer function scope)
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL
	fmt.Printf("Try Slow:\t%q\n", slowURL)
	fmt.Printf("Try Fast:\t%q\n", fastURL)

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
