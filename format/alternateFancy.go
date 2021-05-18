package format

import (
	"math/rand"
	"strings"
)

// AlternateFancy convert to an anternate fancy text
func AlternateFancy(t string) string {
	// Slice of runes of the text
	rs := []rune(t)
	// The new fancy text
	ft := ""

	// Randomly choose if the first rune is upper or lower
	x := rand.Intn(2)

	for i, r := range rs {
		if i%2 == x {
			ft += strings.ToLower(string(r))
		} else {
			ft += strings.ToUpper(string(r))
		}
	}

	return ft
}
