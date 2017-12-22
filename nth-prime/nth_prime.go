// Package prime provides the ability to determine the 'n-th' prime
// number for an integer.
package prime

import "math"

// Nth takes a number n and returns nth prime.
func Nth(n int) (int, bool) {
	var p int
	var count = 0

	for i := 2; count < n; i++ {
		if isPrime(i) == true {
			count++
		}

		if count == n {
			return i, true
		}
	}
	return p, false
}

// IsPrime returns true if a number is prime. We have assumed that
// this is never called with a value of 'n' less than 2.
func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	nsqrt := int(math.Sqrt(float64(n)))

	for i := 3; i < nsqrt+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
