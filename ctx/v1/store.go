package ctx

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ct := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Println("Fetched data without cancel...") // only for tests
			fmt.Fprint(w, d)
		// context has a method Done() which returns a channel which gets sent a signal when the context is "done" or "cancelled"
		case <-ct.Done():
			fmt.Println("Context was done|cancelled...") // only for tests
			store.Cancel()
		}
	}
}
