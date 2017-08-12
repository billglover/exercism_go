package hamming

import (
	"fmt"
)

const testVersion = 6

// Distance computes the Hamming distance between two strings.
// The hamming distance is the number of positions in two strings
// where the characters differ.
// For more detail: [Hamming Distance](https://en.wikipedia.org/wiki/Hamming_distance)
func Distance(a, b string) (d int, err error) {
	if len(a) != len(b) {
		return d, fmt.Errorf("strings are not equal length: %d != %d", len(a), len(b))
	}

	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
