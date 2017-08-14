package house

import "fmt"

const testVersion = 1

var nouns = []string{
	"house",
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

var phrases = []string{
	"Jack built",
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

// Song loops over the 12 verses of the song and concatenates
// them together to form the song
func Song() (s string) {
	for i := 1; i <= 12; i++ {
		if s == "" {
			s = fmt.Sprintf("%s", Verse(i))
			continue
		}
		s = fmt.Sprintf("%s\n\n%s", s, Verse(i))
	}
	return s
}

// Verse takes in a verse index and returns a specific verse
// from the song. The index is 1 based.
func Verse(i int) (v string) {
	return fmt.Sprintf("This is the %s.", Phrase(i-1))
}

// Phrase is a recursive function that returns the phrase
// used to construct a specific verse in the song
func Phrase(i int) (v string) {
	if i == 0 {
		v = fmt.Sprintf("%s that %s", nouns[i], phrases[i])
		return v
	}
	v = fmt.Sprintf("%s\nthat %s the %s", nouns[i], phrases[i], Phrase(i-1))
	return v
}
