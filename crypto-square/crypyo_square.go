package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

// Encode takes a string and returns the square coded cypher text.
func Encode(plain string) (cypher string) {

	plain = strings.ToLower(plain)
	re := regexp.MustCompile("[a-z0-9]")
	chars := re.FindAllString(plain, -1)

	// calculate size of the rectangle
	r := int(math.Floor(math.Sqrt(float64(len(chars)))))
	c := r
	for ; c*r < len(chars); c++ {
	}

	// correct sizing for constraint c-r <= 1
	if c-r > 1 {
		r++
		c--
	}

	res := ""

	// calculate the cypher text
	for x := 0; x < c; x++ {
		for y := 0; y < r; y++ {
			idx := x + y*c

			if idx < len(chars) {
				res += chars[idx]
			}
		}
		res += " "
	}

	cypher = strings.TrimRight(res, " ")
	return cypher
}
