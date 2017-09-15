package foodchain

const testVersion = 3

var pairs = [][]string{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her.", "that wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

// Verse returns a single verse from the song
func Verse(v int) (verse string) {
	verse = "I know an old lady who swallowed a " + pairs[v-1][0] + ".\n"
	verse += pairs[v-1][1]

	insertFlag := false
	for i := v - 1; i >= 1 && i < len(pairs)-1; i-- {
		verse += "\nShe swallowed the " + pairs[i][0] + " to catch the " + pairs[i-1][0]
		if len(pairs[i-1]) == 3 {
			verse += " " + pairs[i-1][2]
		}
		verse += "."
		insertFlag = true
	}

	if insertFlag {
		verse += "\n" + pairs[0][1]
	}

	return verse
}

// Verses returns all verses in the range (vs, ve).
func Verses(vs, ve int) (verses string) {
	for v := vs; v <= ve; v++ {
		verses += Verse(v)

		if v != ve {
			verses += "\n\n"
		}
	}
	return verses
}

// Song returns all verses in the song.
func Song() (song string) {
	song = Verses(1, len(pairs))

	return song
}
