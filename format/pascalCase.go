package format

import (
	"fmt"
	"strings"
)

// PascalCase convert a string to pascal case
func PascalCase(t string) string {
	// Split text into words
	ws := strings.Split(t, " ")
	// The new pascal case text
	pt := ""

	for _, w := range ws {
		rs := []rune(w)

		pt += fmt.Sprintf(
			"%s%s",
			strings.ToUpper(string(rs[0])),
			string(rs[1:]),
		)
	}

	return pt
}
