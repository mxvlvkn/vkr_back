package utils

import (
	"strings"
	"unicode"
)

func CamelCaseToSnake(s string) string {
	var result strings.Builder

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}