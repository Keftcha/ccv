package format

import (
	"fmt"
	"strings"
)

// CamelCase convert a string to camel case
func CamelCase(t string) string {
	// Split text into words
	ws := strings.Split(t, " ")
	// The new pascal case text
	pt := ws[0]

	for i := 1; i < len(ws); i++ {
		rs := []rune(ws[i])

		pt += fmt.Sprintf(
			"%s%s",
			strings.ToUpper(string(rs[0])),
			string(rs[1:]),
		)
	}

	return pt
}
