package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) (bytes int, err error) {
	return fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	_, err := Greet(w, "world")
	if err != nil {
		w.WriteHeader(500)
		_, err := w.Write([]byte("Internal Server Error"))
		if err != nil {
			return
		}
	}
}

// to run: go run .
// open browser at: http://localhost:5001
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
