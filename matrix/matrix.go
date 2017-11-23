// Package matrix takes a string representing a matrix of numbers,
// and provides functions that return the rows and columns of that matrix.
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix
type Matrix [][]int

// New takes a string and parses it into a matrix value. An
// error is returned if it is unable to parse the string.
func New(s string) (Matrix, error) {

	rs := strings.Split(s, "\n")
	m := make([][]int, len(rs))

	for ri, r := range rs {
		cs := strings.Split(strings.TrimSpace(r), " ")
		m[ri] = make([]int, len(cs))

		for ci, c := range cs {
			i, err := strconv.Atoi(strings.TrimSpace(c))
			if err != nil {
				return m, err
			}
			m[ri][ci] = i
		}

		if ri > 0 && len(m[ri]) != len(m[0]) {
			return m, fmt.Errorf("inconsistent number of columns")
		}
	}

	return m, nil
}

// Rows returns a copy of the rows in a Matrix
func (m Matrix) Rows() [][]int {
	nm := make([][]int, len(m))
	for r := range m {
		nm[r] = make([]int, len(m[r]))
		copy(nm[r], (m)[r])
	}
	return nm
}

// Cols returns a copy of the columns in a Matrix
func (m Matrix) Cols() [][]int {
	nm := make([][]int, len(m[0]))
	for ri := range nm {
		nm[ri] = make([]int, len(m))

		for ci := range m {
			nm[ri][ci] = m[ci][ri]
		}
	}

	return nm
}

// Set overwrites an individual element in a Matrix
// with the value provided.
func (m *Matrix) Set(r, c, v int) bool {
	if r > len(*m)-1 || r < 0 {
		return false
	}
	if c > len((*m)[r])-1 || c < 0 {
		return false
	}
	(*m)[r][c] = v
	return true
}
