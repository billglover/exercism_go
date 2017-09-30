package romannumerals

import (
	"fmt"
	"strings"
)

const testVersion = 4

// ToRomanNumeral takes a positive non-zero integer and returns the
// string containing the Roman numberal representation of the number.
// An error is returned if the number passed is a non-positive integer
// or if it is larger than 3000.
func ToRomanNumeral(n int) (roman string, err error) {

	if n <= 0 || n > 3000 {
		return roman, fmt.Errorf("number out of range: %d does not lie in the range 0<n<=3000", n)
	}

	for n > 0 {
		switch {
		case n >= 1000:
			roman += "M"
			n -= 1000
			continue
		case n >= 500:
			roman += "D"
			n -= 500
			continue
		case n >= 100:
			roman += "C"
			n -= 100
			continue
		case n >= 50:
			roman += "L"
			n -= 50
			continue
		case n >= 10:
			roman += "X"
			n -= 10
			continue
		case n >= 5:
			roman += "V"
			n -= 5
			continue
		case n >= 1:
			roman += "I"
			n--
			continue
		default:
			n = 0
			break
		}
	}

	// runs of three or more numerals are not allowed
	roman = strings.Replace(roman, "IIII", "IV", -1)
	roman = strings.Replace(roman, "XXXX", "XL", -1)
	roman = strings.Replace(roman, "CCCC", "CD", -1)

	// we want the shortest possible representation of
	// a given number
	roman = strings.Replace(roman, "VIV", "IX", -1)
	roman = strings.Replace(roman, "LXL", "XC", -1)
	roman = strings.Replace(roman, "DCD", "CM", -1)

	return roman, err
}
