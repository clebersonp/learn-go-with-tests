package v2

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("without cancel signal", func(t *testing.T) {
		data := "hello, world"
		srv := Server(&SpyStore{response: data})

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

		ctx, cancelFunc := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancelFunc)
		request = request.WithContext(ctx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}

	})
}
