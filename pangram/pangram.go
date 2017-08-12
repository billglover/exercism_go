package pangram

import "strings"

const testVersion = 1

// IsPangram returns true if the string contains every
// letter of the alphabet at least once
func IsPangram(s string) bool {
	m := make(map[rune]bool)
	s = strings.ToLower(s)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			m[c] = true
		}
	}

	// the length of our map will indicate the number of unique
	// characters found
	return len(m) == 26
}
