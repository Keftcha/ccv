package format

import (
	"strings"
)

// SnakeCase convert a string to snake case
func SnakeCase(t string, toUpper, toLower bool) string {
	// Replace " " by "_"
	t = strings.ReplaceAll(t, " ", "_")

	if toUpper {
		t = strings.ToUpper(t)
	}

	if toLower {
		t = strings.ToLower(t)
	}

	return t
}
