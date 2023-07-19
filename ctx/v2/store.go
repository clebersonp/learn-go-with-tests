package ct

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := store.Fetch(r.Context())
		if err != nil {
			log.Println("some error occurred:", err)
			return
		}
		fmt.Fprint(w, result)
	}
}
