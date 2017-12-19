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

	// Transpose the input text by swapping rows and cols.
	// We also pad any lines shorter than the max line length.
	opt := make([]string, trows)
	for col := 0; col < tcols; col++ {
		ipt[col] = leftPad(ipt[col], " ", mll-len(ipt[col]))
		for row := 0; row < trows; row++ {
			opt[row] += string(ipt[col][row])
		}
	}

	// Trim trailing space from the final row.
	if trows != 0 {
		opt[trows-1] = strings.TrimRight(opt[trows-1], " ")
	}

	return opt
}

// LeftPad pads a string `s` with `n` times character `p`.
func leftPad(s, p string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += p
	}
	return s + out
}
