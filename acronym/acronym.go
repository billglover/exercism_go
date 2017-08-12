package acronym

import (
	"strings"
)

const testVersion = 3

// Abbreviate takes a string and returns an acronym by
// concatenating and capitalising the first letter of
// each word.
// Valid word separators are: ' ', '-'
func Abbreviate(s string) (a string) {
	w := strings.FieldsFunc(s, split)
	for _, word := range w {
		a += string(word[0])
	}
	return strings.ToUpper(a)
}

func split(r rune) bool {
	switch r {
	case ' ':
		return true
	case '-':
		return true
	default:
		return false
	}
}
