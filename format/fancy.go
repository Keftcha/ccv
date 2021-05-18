package format

import (
	"math/rand"
	"strings"
)

// Fancy convert to a fancy text
func Fancy(t string) string {
	// Slice of runes of the text
	rs := []rune(t)
	// The new fancy text
	ft := ""

	for _, r := range rs {
		if rand.Intn(10)%2 == 0 {
			ft += strings.ToUpper(string(r))
		} else {
			ft += strings.ToLower(string(r))
		}
	}

	return ft
}
