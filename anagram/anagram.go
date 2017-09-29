package anagram

import (
	"strings"
)

const testVersion = 2

// Detect takes a word and identifies anagrams from a list of candidates.
// It returns a slice containing only the valid anagrams.
func Detect(subject string, candidates []string) (anagrams []string) {

	subject = strings.ToLower(subject)

	for _, c := range candidates {
		if isAnagram(subject, strings.ToLower(c)) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

// isAnagram checks if one word is an anagram of another.
// It assumes that both subject and candidate are the same
// case.
func isAnagram(subject, candidate string) bool {

	// if the subject isn't the same length as the word
	// it can't be an anagram
	if len(subject) != len(candidate) {
		return false
	}

	// identical words are not anagrams
	if subject == candidate {
		return false
	}

	for _, r := range subject {
		if strings.Count(subject, string(r)) != strings.Count(candidate, string(r)) {
			return false
		}
	}

	return true
}
