package luhn

import (
	"strconv"
	"strings"
)

const testVersion = 2

// Valid takes a number and validates it using the Luhn check
// algorithm
func Valid(s string) bool {

	s = strings.Trim(s, " ")

	// strings of length 1 or less are not valid
	if len(s) <= 1 {
		return false
	}

	sum := 0
	idx := 1

	for i := len(s); i > 0; i-- {

		// ignore spaces
		if s[i-1:i] == " " {
			continue
		}

		d, err := strconv.Atoi(s[i-1 : i])

		// non-digit characters are not allowed
		if err != nil {
			return false
		}

		// double every second digit
		if idx%2 == 0 {
			d *= 2

			// if doubling results in a number greater than 9 then subtract
			// 9 from the product
			if d >= 9 {
				d -= 9
			}
		}

		sum += d
		idx++
	}

	// if the sum is not evenly divisible by 10 then the number
	// is invalid
	if sum%10 != 0 {
		return false
	}

	return true
}
