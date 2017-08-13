package diffsquares

const testVersion = 1

// SquareOfSums returns the square of the sums of the first
// N numbers.
func SquareOfSums(n int) int {
	s := n * (1 + n) / 2
	return s * s
}

// SumOfSquares returns the sum of the squares of the first
// N numbers.
func SumOfSquares(n int) int {
	// https://proofwiki.org/wiki/Sum_of_Sequence_of_Squares
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the SquareOfSums
// the SumOfSquares of the first N numbers.
func Difference(n int) (d int) {
	return SquareOfSums(n) - SumOfSquares(n)
}
