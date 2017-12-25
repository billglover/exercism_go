// Package beer provides functions to generate the beer song
// If you are unfamiliar with the beer song, the code should
// speak for itself, but more can be found on Wikipedia.
// https://en.wikipedia.org/wiki/99_Bottles_of_Beer
package beer

import (
	"fmt"
	"strconv"
	"unicode"
)

// Song returns the beer song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}

// Verses returns a range of verses from the beer song.
func Verses(v1, v2 int) (string, error) {

	if v1 < v2 {
		return "", fmt.Errorf("invalid range: %d < %d", v1, v2)
	}

	vs := ""
	for v := v1; v >= v2; v-- {
		cv, err := Verse(v)
		if err != nil {
			return vs, err
		}
		vs += cv + "\n"
	}

	return vs, nil
}

// Verse returns a single verse from the beer song.
func Verse(v int) (string, error) {
	if v < 0 || v > 99 {
		return "", fmt.Errorf("invalid verse: %d is not in the range [0,100]", v)
	}

	cur, q := bottles(v)
	next, _ := bottles(v - 1)
	v1 := fmt.Sprintf("%[1]s of beer on the wall, %[1]s of beer.", cur)
	v2 := fmt.Sprintf("Take %s down and pass it around, %s of beer on the wall.", q, next)

	// The last verse is a special case so handle this as an exception.
	if v == 0 {
		v2 = fmt.Sprintf("Go to the store and buy some more, %s of beer on the wall.", next)
	}

	// Capitalise the first letter of the verse. We use `unicode.ToUpper` as we don't
	// make any assumptions around the character set used in the verse.
	v1r := []rune(v1)
	v1r[0] = unicode.ToUpper(v1r[0])
	v1 = string(v1r)

	// Return the verse along with a trailing new-line character.
	return v1 + "\n" + v2 + "\n", nil
}

// Bottles is a helper function that takes a number of bottles and returns the correctly
// pluralised form along with a quantification string.
func bottles(b int) (cur, q string) {
	switch b {
	case 1:
		cur = "1 bottle"
		q = "it"
	case 0:
		cur = "no more bottles"
	case -1:
		cur = "99 bottles"
	default:
		cur = strconv.Itoa(b) + " bottles"
		q = "one"
	}
	return cur, q
}
