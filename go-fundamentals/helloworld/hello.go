package main

import "fmt"

// in terminal: go run .    or     go run hello.go

func hello(name string) string {
	return "Hello, " + name // -> the "domain"
}

func main() {
	// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
	// separated the side effects from "domain" part. So it's easier to test
	fmt.Println(hello("world"))
}