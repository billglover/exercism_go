package diamond

import "fmt"

const testVersion = 1

// Gen takes a letter (byte), and returns a string representing a diamond
// shape. The string returned starts with 'A', and contains the supplied
//  letter at the widest point.
func Gen(l byte) (d string, err error) {

	// we only handle letters in the range [A-Z]
	if l < byte('A') || l > byte('Z') {
		return d, fmt.Errorf("character out of range: want [A-Z], got %s", string(l))
	}

	rows := int(l - byte('A'))

	// populate the top half of the diamond
	for r := 0; r <= rows; r++ {
		d += addRow(r, rows)
	}

	// if our diamond has only one row we cna return
	if rows == 0 {
		return d, err
	}

	// populate the bottom half of the diamond
	for r := rows - 1; r >= 0; r-- {
		d += addRow(r, rows)
	}

	return d, err
}

// addRow takes a row index and total number of rows
// and returns the string representation of the row
// from the diamond.
func addRow(r, rows int) (s string) {
	l := byte(r) + byte('A')

	chars := make([]rune, (rows*2)+1)
	for ci := range chars {
		chars[ci] = ' '
	}

	oPad := rows - r
	iPad := (r * 2) - 1

	chars[oPad] = rune(l)
	chars[oPad+iPad+1] = rune(l)
	s = string(chars)

	s += "\n"

	return s
}
