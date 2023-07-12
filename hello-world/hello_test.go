package main

import "testing"

// Writing tests
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T

// Install godoc to run locally
// $ go install golang.org/x/tools/cmd/godoc@latest
// Run godoc
// $ godoc -http :8000
// http://localhost:8000/pkg/

// Run tests
// $ go test


// TDD cycle
// Write a test
// Make the compiler pass
// Run the test, see that it fails and check the error message is meaningful
// Write enough code to make the test pass
// Refactor

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say, 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Frank", "French")
		want := "Bonjour, Frank"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	
	// TB interface that T and B satisfied
	// show actual test line if fail
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}
