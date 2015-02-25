package stringcase

import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"

// ToPropertyCase converts the string to 'property-case' and returns it.
func ToPropertyCase(s string) string {

	// Convert.
	result := strings.Map(
		func(r rune) rune {

			if whitespace.IsWhitespace(r) || '_' == r {
				return '-'
			} else {
				return unicode.ToLower(r)
			}
		},
		s)

	// Return
	return result
}

// FromPropertyCase converts the "property-case" string to a spaced string
// "property case" and returns it.
func FromPropertyCase(s string) string {
	return strings.Replace(s, "-", " ", -1)
}
