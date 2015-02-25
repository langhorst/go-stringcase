package stringcase

import "github.com/reiver/go-whitespace"
import "strings"
import "unicode"

// ToConstCase converts the string to 'CONST_CASE' and returns it.
func ToConstCase(s string) string {

	// Convert.
	result := strings.Map(
		func(r rune) rune {

			if whitespace.IsWhitespace(r) || '-' == r {
				return '_'
			} else {
				return unicode.ToUpper(r)
			}
		},
		s)

	// Return
	return result
}

// FromConstCase converts the 'CONST_CASE' string to a spaced string and
// returns it.
func FromConstCase(s string) string {
	return strings.Replace(s, "_", " ", -1)
}
