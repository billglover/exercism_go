package atbash

import (
	"strings"
)

var cipher = []byte("zyxwvutsrqponmlkjihgfedcba")

// Atbash takes a plain text string and returns the Atbash
// cipher text equivalent.
func Atbash(pt string) string {

	pt = strings.ToLower(pt)
	pb := []byte(pt)

	// We make the assumption that the maximum length of the cipher
	// text is the length of the plain text with some additional
	// characters for padding.
	maxlen := len(pt) + len(pt)/5 + 1
	c := make([]byte, maxlen)
	ci := 0
	pad := 0

	for _, pchar := range pb {

		// If we have a digit, pass it straight through adding
		// padding if necessary.
		if pchar >= '0' && pchar <= '9' {
			if ci%5 == 0 && ci > 0 {
				c[ci+pad] = ' '
				pad++
			}
			c[ci+pad] = pchar
			ci++
			continue
		}

		// If we have a character, look-up the ciphertext
		// equivalent adding padding if necessary.
		if pchar >= 'a' && pchar <= 'z' {
			if ci%5 == 0 && ci > 0 {
				c[ci+pad] = ' '
				pad++
			}
			c[ci+pad] = cipher[pchar-'a']
			ci++
			continue
		}

	}

	return string(c[:ci+pad])
}
