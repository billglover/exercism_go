package diffsquares

const testVersion = 1

// SquareOfSums returns the square of the sums of the first
// N numbers.
func SquareOfSums(n int) (s int) {
	for ; n > 0; n-- {
		s += n
	}
	return s * s
}

// SumOfSquares returns the sum of the squares of the first
// N numbers.
func SumOfSquares(n int) (s int) {
	for ; n > 0; n-- {
		s += (n * n)
	}
	return s
}

// Difference returns the difference between the SquareOfSums
// the SumOfSquares of the first N numbers.
func Difference(n int) (d int) {
	return SquareOfSums(n) - SumOfSquares(n)
}
