package grains

import (
	"fmt"
)

const testVersion = 1

// Total returns the total number of grains on a chess board if the
// number of grains on each square doubles
func Total() (g uint64) {
	for n := 1; n <= 64; n++ {
		sg, _ := Square(n)
		g += sg
	}
	return g
}

// Square returns the total number of grains on a square of a chess
// board, given that the number of grains on each square is double
// that of the previous square
func Square(n int) (g uint64, err error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("invalid square: %d is not in the expected range {1, 64}", n)
	}

	if n == 1 {
		return 1, nil
	}

	g, err = Square(n - 1)
	return 2 * g, err
}
