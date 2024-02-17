package main

import "testing"

// in terminal: go test   or more verbose    go test -v

// Writing a test is just like writing a function, with a few rules
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

func TestHello(t *testing.T) {
	got := hello("Chris")
	want := "Hello, Chris"

	if got != want {
		// For tests %q is very useful as it wraps your values in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
