package brackets

import (
	"strings"
)

const testVersion = 5

// Bracket takes a string containing brackets `[]`, braces `{}`
// and parentheses `()` and verifies that all the pairs are matched
// and nested correctly.
func Bracket(s string) (bool, error) {

	s = removeNonBrackets(s)
	lastLen := len(s) + 1
	for len(s) < lastLen {
		lastLen = len(s)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)
	}

	return s == "", nil
}

// removeNonBrackets removes all non brackets, braces
// and parentheses from a string.
func removeNonBrackets(s string) string {
	for _, r := range s {
		if isBracket(r) == false {
			s = strings.Replace(s, string(r), "", -1)
		}
	}
	return s
}

// isBracket validates whether a rune is a a bracket,
// brace or parentheses
func isBracket(r rune) bool {
	switch r {
	case '{',
		'}',
		'[',
		']',
		'(',
		')':
		return true
	}
	return false
}
