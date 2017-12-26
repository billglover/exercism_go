// Package beer provides functions to generate the beer song
// If you are unfamiliar with the beer song, the code should
// speak for itself, but more can be found on Wikipedia.
// https://en.wikipedia.org/wiki/99_Bottles_of_Beer
package beer

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Song returns the beer song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}

// Verses returns a range of verses from the beer song.
func Verses(v1, v2 int) (string, error) {

	if v1 < v2 || v1 > 99 || v2 < 0 {
		return "", fmt.Errorf("invalid range: %d < %d", v1, v2)
	}

	var buffer bytes.Buffer
	for v := v1; v >= v2; v-- {
		cv, err := Verse(v)
		if err != nil {
			return "", err
		}
		buffer.WriteString(cv)
		buffer.WriteString("\n")
	}

	return buffer.String(), nil
}

// Verse returns a single verse from the beer song.
func Verse(v int) (string, error) {
	if v < 0 || v > 99 {
		return "", fmt.Errorf("invalid verse: %d is not in the range [0,100]", v)
	}

	var buffer bytes.Buffer

	// generate the first line of the verse
	b := bottles(v)
	buffer.WriteString(strings.ToUpper(string(b[0])))
	buffer.WriteString(b[1:])
	buffer.WriteString(" of beer on the wall, ")
	buffer.WriteString(b)
	buffer.WriteString(" of beer.\n")

	// generate the second line of the verse
	if v == 0 {
		buffer.WriteString("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
		return buffer.String(), nil
	}

	q := "one"
	if v == 1 {
		q = "it"
	}
	b = bottles(v - 1)

	buffer.WriteString("Take ")
	buffer.WriteString(q)
	buffer.WriteString(" down and pass it around, ")
	buffer.WriteString(b)
	buffer.WriteString(" of beer on the wall.\n")

	return buffer.String(), nil
}

// Bottles is a helper function that takes a number of bottles and returns the correctly
// pluralised form along with a quantification string.
func bottles(b int) string {
	if b == 1 {
		return "1 bottle"
	}
	if b == 0 {
		return "no more bottles"
	}

	return strconv.Itoa(b) + "bottles"
}
