package palindrome

import (
	"fmt"
)

const testVersion = 1

// Product holds a slice of factors for a given product
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products returns all the palindrome products in a given range.
func Products(min, max int) (pMin, pMax Product, err error) {
	if min > max {
		err := fmt.Errorf("fmin > fmax")
		return pMin, pMax, err
	}

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			p := a * b
			if isPalindrome(p) {
				if p == pMax.Product {
					pMax.Factorizations = append(pMax.Factorizations, [2]int{a, b})
					continue
				}
				if p > pMax.Product {
					pMax.Product = p
					pMax.Factorizations = [][2]int{{a, b}}
				}
				if pMin.Product == 0 {
					pMin.Product = p
					pMin.Factorizations = [][2]int{{a, b}}
				}
			}
		}
	}

	if pMin.Product == 0 && pMax.Product == 0 {
		err = fmt.Errorf("no palindromes")
	}

	return pMin, pMax, err
}

func isPalindrome(a int) bool {
	s := fmt.Sprintf("%d", a)
	rs := reverseString(s)
	if s == rs {
		return true
	}
	return false
}

func reverseString(s string) string {
	rs := ""
	for c := len(s) - 1; c >= 0; c-- {
		rs += string(s[c])
	}
	return rs
}
