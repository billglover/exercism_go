// Package igpay provides functions to translate between English and Pig Latin.
package igpay

import (
	"strings"
)

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}

// PigLatin translates between English and Pig Latin.
func PigLatin(s string) string {
	ws := strings.Split(s, " ")
	for wi := range ws {
		ws[wi] = encodeWord(ws[wi])
	}
	return strings.Join(ws, " ")
}

// EncodeWord translates a single word from English to Pig Latin
func encodeWord(s string) string {
	en := []byte(s)
	pl := make([]byte, len(en)+2)

	// determine whether the initial character is a vowel of consonant
	initialValue := isVowel(en, 0)
	index := 1

	// scan to find the end of the initial run of consonants or vowels
	for index < len(en) {
		if isVowel(en, index) != initialValue {
			break
		}
		index++
	}

	// encode the PigLatin version of the words
	if initialValue {
		pl = en
	} else {
		pl = en[index:]
		pl = append(pl, en[0:index]...)
	}
	pl = append(pl, []byte{'a', 'y'}...)

	return string(pl)
}

// IsVowel returns true if a character is a vowel. It handles a couple of exceptions
// that are specific to PigLatin, e.g. 'xray', 'yttria', 'queen', etc.
func isVowel(b []byte, pos int) bool {

	// 'u' preceded by 'q' should be treated as a consonant
	if b[pos] == 'u' && b[pos-1] == 'q' {
		return false
	}

	// check for the usual vowels
	for i := range vowels {
		if b[pos] == vowels[i] {
			return true
		}
	}

	// 'y' or 'x' followed by a consonant should be treated as a vowel
	if b[pos] == 'y' || b[pos] == 'x' {
		for i := range vowels {
			if b[pos+1] == vowels[i] {
				return false
			}
		}
		return true
	}

	// everything else is a constant
	return false
}
