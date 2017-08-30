package pythagorean

import (
	"math"
)

const testVersion = 1

// Triplet represents a Pythagorean triplet, made up of 3 integers
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (ts []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if c == math.Trunc(c) && int(c) <= max {
				ts = append(ts, Triplet{a, b, int(c)})
			}
		}
	}

	return ts
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) (ts []Triplet) {
	pts := Range(1, p)
	for _, t := range pts {
		if t[0]+t[1]+t[2] == p {
			ts = append(ts, t)
		}
	}

	return ts
}
