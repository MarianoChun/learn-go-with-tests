package iteration

// Repeat takes a character and a repeat count and return a string with that same character repeated the specified number of repetitions
func Repeat(character string, repeatCount int) string {
	var result string
	for i := 0; i < repeatCount; i++ {
		result += character
	}
	return result
}
