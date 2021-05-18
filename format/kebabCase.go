package format

import (
	"strings"
)

// KebabCase convert a string to kebab case
func KebabCase(t string, toUpper, toLower bool) string {
	// Replace " " by "-"
	t = strings.ReplaceAll(t, " ", "-")

	if toUpper {
		t = strings.ToUpper(t)
	}

	if toLower {
		t = strings.ToLower(t)
	}

	return t
}
