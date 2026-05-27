package function

import "unicode"

func ToLowerCase(text string) string {
	char := []rune(text)
	for i := range char {
		char[i] = unicode.ToLower(char[i])
	}
	return string(char)
}
