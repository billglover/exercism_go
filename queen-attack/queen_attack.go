package queenattack

import (
	"fmt"
)

const testVersion = 2

// CanQueenAttack takes the position of two queens on a
// chess board and returns true if they are able to attack
// each other.
func CanQueenAttack(q1, q2 string) (attac bool, err error) {

	a, err := point(q1)
	if err != nil {
		return false, err
	}

	b, err := point(q2)
	if err != nil {
		return false, err
	}

	if a == b {
		return false, fmt.Errorf("queens share same location")
	}

	if onStraight(a, b) == true {
		return true, nil
	}

	if onDiagonal(a, b) == true {
		return true, nil
	}

	return false, nil
}

func point(l string) (p [2]int, err error) {
	if len(l) != 2 {
		return p, fmt.Errorf("invalid point (%s)", l)
	}

	p1 := int(l[0]) - 96
	if p1 < 1 || p1 > 8 {
		return p, fmt.Errorf("invalid point (%s)", l)
	}
	p[0] = p1

	p2 := int(l[1]) - 48
	if p2 < 1 || p2 > 8 {
		return p, fmt.Errorf("invalid point (%s)", l)
	}
	p[1] = p2

	return p, nil
}

func onStraight(a, b [2]int) bool {
	if a[0] == b[0] || a[1] == b[1] {
		return true
	}
	return false
}

func onDiagonal(a, b [2]int) bool {
	g := (float32(a[0]-b[0]) / float32(a[1]-b[1]))
	if g*g == 1 {
		return true
	}
	return false
}
