package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		// start a test server to simulate a response with delay
		slowServer := makeDelayedServer(5 * time.Millisecond)
		// start a test server to simulate the faster response without delay
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// defer to close servers at the end of performing this block of code (TestRacer function scope)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		// start a test server to simulate a response with delay
		serverA := makeDelayedServer(11 * time.Second)
		// start a test server to simulate the faster response without delay
		serverB := makeDelayedServer(12 * time.Second)

		// defer to close servers at the end of performing this block of code (TestRacer function scope)
		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
