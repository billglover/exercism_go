package bob

import "strings"

const testVersion = 3

// Hey allows you to communicate with Bob by passing a string
// as the input. Bob responds by returning a string.
// Bob answers 'Sure.' if you ask him a question.
// He answers 'Whoa, chill out!' if you yell at him.
// He says 'Fine. Be that way!' if you address him without actually saying
// anything.
// He answers 'Whatever.' to anything else.
func Hey(s string) (r string) {

	s = strings.TrimSpace(s)

	if s == "" {
		return "Fine. Be that way!"
	}

	shouting := true
	charsFound := false
	for _, c := range s {

		if c >= 'a' && c <= 'z' {
			shouting = false
			charsFound = true
		}
		if c >= 'A' && c <= 'Z' {
			charsFound = true
		}
	}

	if shouting == true && charsFound == true {
		return "Whoa, chill out!"
	}

	if s[len(s)-1] == '?' {
		return "Sure."
	}

	return "Whatever."
}
