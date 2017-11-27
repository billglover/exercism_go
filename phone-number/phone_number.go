// Package phonenumber provides functions for cleaning up user entered phone
// numbers so that they can be represented as valid NANP region numbers.
// NANP numbers are ten-digit numbers consisting of a three-digit Numbering
// Plan Area code, commonly known as *area code*, followed by a seven-digit
// local number. The first three digits of the local number represent the
// *exchange code*, followed by the unique four-digit number which is the
// *subscriber number*.
package phonenumber

import (
	"fmt"
)

// Number returns a string representing the numeric representation of a
// phone number.
func Number(s string) (string, error) {
	cs := clean(s)

	// strip the country code if present
	if cs[0] == '1' {
		cs = cs[1:]
	}

	// return an error if the number is not 10 digits
	if len(cs) != 10 {
		return cs, fmt.Errorf("invalid length")
	}

	// area codes must begin with a number in the range [2-9]
	if cs[0] < '2' || cs[0] > '9' {
		return cs, fmt.Errorf("invalid area code")
	}

	// exchange codes must begin with a number in the range [2-9]
	if cs[3] < '2' || cs[3] > '9' {
		return cs, fmt.Errorf("invalid exchange code")
	}

	return cs, nil
}

// AreaCode returns a string representing the area code of a phone number.
func AreaCode(s string) (string, error) {
	return s, nil
}

// Format produces a human readable representation of a phone number.
func Format(s string) (string, error) {
	return s, nil
}

func clean(s string) string {
	count := 0
	c := make([]byte, len(s))
	for _, b := range []byte(s) {
		if b >= '0' && b <= '9' {
			c[count] = b
			count++
		}
	}
	return string(c[:count])
}
