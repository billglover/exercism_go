// Package prime provides a function to return the prime factors
// of a number in increasing order.
package prime

// Factors returns prime factors in increasing order
func Factors(i int64) []int64 {
	r := []int64{}

	for {
		if i%2 != 0 {
			break
		}
		i = i / 2
		r = append(r, 2)
	}

	for o := int64(0); o <= i; {
		if i%(3+o) != 0 {
			o += 2
			continue
		}
		i = i / (3 + o)
		r = append(r, 3+o)
	}

	return r
}
