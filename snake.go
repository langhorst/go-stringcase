package stringcase

import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"

// ToSnakeCase converts the string to 'snake_case' and returns it.
func ToSnakeCase(s string) string {

	// Convert.
	result := strings.Map(
		func(r rune) rune {

			if whitespace.IsWhitespace(r) || '-' == r {
				return '_'
			} else {
				return unicode.ToLower(r)
			}
		},
		s)

	// Return
	return result
}

// FromSnakeCase converts the "snake_case" string to a spaced string
// "snake case" and returns it.
func FromSnakeCase(s string) string {
	return strings.Replace(s, "_", " ", -1)
}
