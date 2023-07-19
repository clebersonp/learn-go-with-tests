package ct

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

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled...")
				return
			default:
				time.Sleep(15 * time.Millisecond)
				char := string(c)
				result += char
				log.Printf("concat %q to result...\n", char) // only to debug
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		log.Println("context was done|cancelled....") // only to debug
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// We need to test that we do not write any kind of response on the error case.
// implements http.ResponseWriter
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	fmt.Println("Header was performed") // just to debug
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	fmt.Println("Write was performed") // just to debug
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	fmt.Println("WriteHeader was performed") // just to debug
	s.written = true
}

func TestServer(t *testing.T) {
	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", but want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(1*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
