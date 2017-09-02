package isogram

import "unicode"

const testVersion = 1

// IsIsogram determines if a word or phrase is an isogram. The comparison is
// case insensitive and ignores non letter characters.
func IsIsogram(s string) bool {
	a := []rune(s)
	m := make(map[rune]bool, len(s))

	for _, l := range a {
		r := unicode.ToLower(l)
		if unicode.IsLetter(r) && m[r] {
			return false
		}
		m[r] = true
	}

	return true
}
