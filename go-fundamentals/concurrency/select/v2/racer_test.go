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
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(15 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer slowServer.Close()

	// start a test server to simulate the faster response without delay
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
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
