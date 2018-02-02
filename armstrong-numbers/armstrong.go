package armstrong

// IsNumber returns true if a number is an armstrong number
func IsNumber(n int) bool {

	// Single digit numbers are armstrong numbers.
	if n < 10 {
		return true
	}

	sum := 0
	i := n

	l := 0
	for i > 0 {
		l++
		i /= 10
	}
	i = n

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
