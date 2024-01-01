package main

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	italian            = "Italian"
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
	italianHelloPrefix = "Ciao"
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
	case italian:
		prefix = italianHelloPrefix
	}

	return fmt.Sprintf("%s %s!", prefix, name)
}
