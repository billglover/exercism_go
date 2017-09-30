package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Frequency contains a map of words and their
// associated frequencies.
type Frequency map[string]int

// WordCount takes a string and returns the count
// of each word in the string.
func WordCount(phrase string) Frequency {

	// ignore case
	phrase = strings.ToLower(phrase)

	f := make(map[string]int, 1)

	var start int
	var pos int
	var c rune

	inWord := false

	for pos, c = range phrase {

		if unicode.IsLetter(c) || // letters are considered part of a word
			unicode.IsNumber(c) || // as are numbers
			(c == '\'' && // as is the apostrophe
				unicode.IsLetter(rune(phrase[pos-1])) &&
				unicode.IsLetter(rune(phrase[pos+1]))) {

			// if we aren't already in a word, capture the start position
			if inWord == false {
				start = pos
			}

			// we are now inside a word
			inWord = true

		} else {

			// if we have come to the end of a word, increment the counter
			if inWord == true {
				word := phrase[start:pos]
				f[word]++
			}
			inWord = false
		}
	}

	// if we were in a word when we came to the end of the phrase
	// we should capture it
	if inWord == true {
		word := phrase[start:len(phrase)]
		f[word]++
	}

	return f
}
