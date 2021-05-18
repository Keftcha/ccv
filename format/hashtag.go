package format

import (
	"strings"
)

// Hashtag convert a string with hashtag
func Hashtag(t string) string {
	// Split text into wors slice
	ws := strings.Split(t, " ")
	// New words
	nws := make([]string, len(ws))

	for i, w := range ws {
		if !strings.ContainsAny(w, ",\"';:/\\`|#()[]{}.?!<>") && 4 <= len(w) {
			nws[i] = "#" + w
		} else {
			nws[i] = w
		}
	}

	return strings.Join(nws, " ")
}
