package format

import (
	"strings"
)

// Titleize uppercase the first letter of each word
func Titleize(t string) string {
	// Split the text in words
	ws := strings.Split(t, " ")
	// The new titleized text
	tt := make([]string, len(ws))

	for i, w := range ws {
		tt[i] = Capitalize(w)
	}

	return strings.Join(tt, " ")
}
