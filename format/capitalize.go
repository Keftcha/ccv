package format

import (
	"strings"
)

// Capitalize uppercase the first letter
func Capitalize(t string) string {
	// Slice of runes
	rs := []rune(t)

	return strings.ToUpper(string(rs[0])) + string(rs[1:])
}
