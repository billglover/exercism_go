// Package transpose provides a function for transposing
// the text in a slice of strings. Rows become columns
// and columns become rows.
package transpose

import (
	"strings"
)

// Transpose takes a slice of strings and transposes
// each character. If the rows in the input array are
// of different lengths, they are padded to the left.
func Transpose(ipt []string) []string {

	// Find the maximum line length to allow us to pad
	// lines that are shorter than the maximum.
	mll := 0
	for r := range ipt {
		if len(ipt[r]) > mll {
			mll = len(ipt[r])
		}
	}

	// Determine the size of our output slice.
	trows := mll
	tcols := len(ipt)

	// Padding our input strings on the right is
	// equivalent to left padding our output strings.
	for ri := range ipt {
		if len(ipt[ri]) != mll {
			ipt[ri] = rightPad(ipt[ri], ' ', mll-len(ipt[ri]))
		}
	}

	// Transpose the input text by swapping rows and cols.
	// We also pad any lines shorter than the max line length.
	out := make([]string, trows)
	row := make([]byte, tcols)
	for ri := 0; ri < trows; ri++ {
		for ci := 0; ci < tcols; ci++ {
			row[ci] = ipt[ci][ri]
		}
		out[ri] = string(row)
	}

	// Trim trailing space from the final row.
	if trows != 0 {
		out[trows-1] = strings.TrimRight(out[trows-1], " ")
	}

	return out
}

// RightPad pads a string `s` with `n` times character `p`.
func rightPad(s string, p byte, n int) string {
	pad := make([]byte, n)
	for i := range pad {
		pad[i] = p
	}
	return s + string(pad)
}
