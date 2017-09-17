package lsproduct

import (
	"fmt"
	"strconv"
)

const testVersion = 5

// LargestSeriesProduct calculates the largest product for a contiguous
// substring of digits of length r. It returns an error if the string
// contains non-numeric characters or if the length of the contiguous
// substring is invalid. The length of the continguous substring is
// considered to be invalid if it is negative or longer than the string
// provided.
func LargestSeriesProduct(s string, r int) (p int, err error) {

	if r > len(s) {
		return p, fmt.Errorf("series length %d is longer than string length %d", r, len(s))
	}

	if r < 0 {
		return p, fmt.Errorf("invalid series length %d < 0", r)
	}

	for pos := 0; pos <= len(s)-r; pos++ {

		prod := 1

		for c := pos; c < pos+r; c++ {
			i, err := strconv.Atoi(string(s[c]))
			if err != nil {
				return p, err
			}

			prod *= i
		}

		if prod > p {
			p = prod
		}
	}

	return p, err
}
