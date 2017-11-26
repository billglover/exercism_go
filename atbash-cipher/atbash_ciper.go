package atbash

import (
	"strings"
)

var plain = []byte("abcdefghijklmnopqrstuvwxyz")
var cipher = []byte("zyxwvutsrqponmlkjihgfedcba")

// Numbers are passed through
// Capital letters are converted to lowercase
// Punctuation (incl. spaces) is ignored

// Every fifth character should insert a space in output

func Atbash(p string) string {

	p = strings.ToLower(p)

	pi := 0
	ci := 0
	pad := 0

	maxlen := len(p) + len(p)/5 + 1

	c := make([]byte, maxlen)

	for pi < len(p) {

		pchar := p[pi]
		switch {
		case pchar >= '0' && pchar <= '9':
			if ci%5 == 0 && ci > 0 {
				c[ci+pad] = ' '
				pad++
			}
			c[ci+pad] = pchar
			ci++
		case pchar >= 'a' && pchar <= 'z':
			if ci%5 == 0 && ci > 0 {
				c[ci+pad] = ' '
				pad++
			}
			c[ci+pad] = cipher[pchar-'a']
			ci++
		default:
			pi++
			continue
		}

		pi++
	}

	return string(c[:ci+pad])
}
