package iteration

import "strings"

// Repeat takes a character and a repeat count and return a string with that same character repeated the specified number of repetitions
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
