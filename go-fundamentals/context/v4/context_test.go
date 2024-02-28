package v4

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// How to test HTTP handler that has had the request cancelled by the client
// How to use context to manage cancellation
// How to write a function tha accepts context and uses it to cancel itself by using goroutines, select and channels
// Follow Google's guidelines as to how to manage cancellation by propagating request scoped context through your call-stack
// How to roll your own spy for http.ResponseWriter if you need it

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprint(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			// will stop the work if a signal is sent in this channel
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	// finally this select is waiting for that goroutine to finish its work or for the cancellation to occur.
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// we need to test that we do not write any kind of response on the error case.
// httptest.ResponseRecorder doesn't have a way of figuring this out, so we'll have to roll our own spy to test for this
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		resp := response.Body.String()
		if resp != data {
			t.Errorf("got %q, want %q", resp, data)
		}

	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// One of the main points of context is that it is a consistent way of offering cancellation.
		ctx, cancelFunc := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancelFunc)
		request = request.WithContext(ctx)

		response := &SpyResponseWriter{}

		srv.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not hava been written")
		}
	})
}
