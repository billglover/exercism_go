package triangle

import (
	"math"
)

const testVersion = 3

// KindFromSides determines the type of a triangle based on the
// length of its three sides.
func KindFromSides(a, b, c float64) Kind {

	// NaN or Inf both can't be triangles
	if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return NaT
	}

	// If any of the sides is <= 0 it can't be a triangle
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// the sum of any two sides can't be less than the third
	if a+b < c || b+c < a || c+a < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || c == a {
		return Iso
	}

	return Sca
}

// Kind is an integer that indicates a type of triangle
type Kind int

const (
	// NaT = Not a triangle
	NaT = iota

	// Equ = Equilateral triangle
	Equ

	// Iso = Isoceles triangle
	Iso

	// Sca = Scalene triangle
	Sca
)
