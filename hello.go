package main

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

func main() {
	fmt.Println(Hello("World", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return fmt.Sprintf("%s %s!", prefix, name)
}
