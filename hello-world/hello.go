package main

import (
	"fmt"
	"strings"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "spanish"
	french             = "french"
	world              = "World"
)

// Hello function only returns the phrase: "Hello, someName"
func Hello(name, language string) string {
	if name == "" {
		name = world
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch strings.ToLower(language) {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
