package main

import "fmt"

// in terminal: go run .    or     go run hello.go

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	world              = "World"
	spanish            = "Spanish"
	french             = "French"
	emptyStr           = ""
)

func hello(name, language string) string {
	if name == emptyStr {
		name = world
	}
	if language == spanish {
		return fmt.Sprint(spanishHelloPrefix, name)
	}
	if language == french {
		return fmt.Sprint(frenchHelloPrefix, name)
	}
	return fmt.Sprint(englishHelloPrefix, name) // -> the "domain"
}

func main() {
	// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
	// separated the side effects from "domain" part. So it's easier to test
	fmt.Println(hello("world", ""))
}
