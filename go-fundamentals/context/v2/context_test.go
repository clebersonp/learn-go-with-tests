package v2

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, q *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello, world"
	srv := Server(&StubStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	srv.ServeHTTP(response, request)

	resp := response.Body.String()
	if resp != data {
		t.Errorf("got %q, want %q", resp, data)
	}
}
