package say

import "strings"

const testVersion = 1

// Notes:
// - we speak in groups of hundred
// - the first 100 numbers contain some special names
var digits = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

var powersOfTen = map[int]string{
	1:  "ten",
	2:  "hundred",
	3:  "thousand",
	4:  "ten thousand",
	5:  "hundred thousand",
	6:  "million",
	9:  "billion",
	12: "trillion",
	15: "quadrillion",
	18: "quintillion",
}

// Say takes a uint64 and returns a string representing the spoken
// english representation of the number.
func Say(n uint64) (s string) {

	// treat zero as a special case
	if n == 0 {
		return "zero"
	}

	// We speak numbers in 100s of units and so we'll loop through
	// our large number in decreasing powers of 10 counting the
	// number of units we have.
	powers := make([]int, 19)
	for p := 18; p >= 0; {
		pow := intPow(10, p)
		if n >= pow {
			powers[p]++
			n -= pow
			continue
		}
		p -= 3
	}

	// For each of the powers of ten, we construct a phrase that
	// represents the pronounciation of the number of units.
	for p := len(powers) - 1; p >= 0; p-- {
		if powers[p] > 0 {
			s += sayHundreds(powers[p]) + " " + powersOfTen[p] + " "
		}
	}

	// Tidy things up by trimming the white space from the
	// phrase representing the number.
	return strings.TrimSpace(s)
}

// intPow is a helper function to calculate a**b when a and b
// are both integers.
func intPow(a, b int) uint64 {
	r := uint64(1)
	for b > 0 {
		r *= uint64(a)
		b--
	}
	return r
}

func sayHundreds(n int) string {
	phrase := ""

	if phrase, ok := digits[n]; ok == true {
		return phrase
	}

	// figure out the 100s
	h := n / 100
	if s, ok := digits[h]; ok == true {
		phrase += s + " " + powersOfTen[2] + " "
	}
	n = n % 100

	// figure out the tens
	if s, ok := digits[n]; ok == true {
		phrase += s
		return strings.TrimSpace(phrase)
	}

	t := n / 10
	u := n % 10
	if s, ok := digits[t*10]; ok == true {
		if u == 0 {
			phrase += s
			return strings.TrimSpace(phrase)
		}

		phrase += s + "-"
	}

	// figure out the units
	if s, ok := digits[u]; ok == true {
		phrase += s + " "
	}

	// tidy things up before returning the string
	return strings.TrimSpace(phrase)
}
