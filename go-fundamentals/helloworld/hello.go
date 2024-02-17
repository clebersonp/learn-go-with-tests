package main

import "fmt"

// in terminal: go run .    or     go run hello.go

const (
	// utilities
	emptyStr = ""
	world    = "World"

	// languages
	spanish    = "Spanish"
	french     = "French"
	portuguese = "Portuguese"

	// prefixes
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

func hello(name, language string) (greetings string) {
	if name == emptyStr {
		name = world
	}
	return fmt.Sprint(greetingPrefix(language), name) // -> the "domain"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // will return the prefix value declared as 'named return value'
}

func main() {
	// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
	// separated the side effects from "domain" part. So it's easier to test
	fmt.Println(hello("world", ""))
}
