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
	if language == spanish {
		return fmt.Sprintf("%s %s!", spanishHelloPrefix, name)
	}
	if language == french {
		return fmt.Sprintf("%s %s!", frenchHelloPrefix, name)
	}
	return fmt.Sprintf("%s %s!", englishHelloPrefix, name)
}
