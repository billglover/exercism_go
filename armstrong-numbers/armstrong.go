package armstrong

import "strconv"

// IsNumber returns true if a number is an armstrong number
func IsNumber(n int) bool {

	// Single digit numbers are armstrong numbers.
	if n < 10 {
		return true
	}

	l := len(strconv.Itoa(n))
	sum := 0

	i := n
	for i > 0 {
		sum += exp(i%10, l)
		i /= 10
	}
	return sum == n
}

func exp(a, b int) int {
	result := 1

	for b != 0 {
		if (b & 1) != 0 {
			result *= a
		}
		b >>= 1
		a *= a
	}

	return result
}
