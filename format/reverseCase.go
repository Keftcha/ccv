package format

import (
	"strings"
	"unicode"
)

// ReverseCase reverse the case of a text
func ReverseCase(t string) string {
	// Slice of runes of the text
	rs := []rune(t)
	// The new text
	t = ""

	for _, r := range rs {
		if unicode.IsUpper(r) {
			t += strings.ToLower(string(r))
		} else if unicode.IsLower(r) {
			t += strings.ToUpper(string(r))
		} else {
			t += string(r)
		}
	}

	return t
}
