package robotname

import (
	"math/rand"
)

const testVersion = 1

// NameRegister is a register of names.
type NameRegister map[string]bool

// robotRegister holds the name of every robot
// ever issued.
var robotRegister = NameRegister{}

// Robot represents a robot
type Robot struct {
	name string
}

// Name returns the name of a robot. If the robot
// hasn't been issued a name yet, it gets given
// a unique name.
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = newName()
	}
	return r.name
}

// Reset erases a robots memory and causes it to
// forget its name.
func (r *Robot) Reset() {
	r.name = ""
}

// newName generates a name for a robot and ensures
// that the name is unique. The name takes the format
// of two uppercase letters followed by three digits.
func newName() (n string) {

	for unique := false; unique == false; {

		chars := make([]rune, 5)
		chars[0] = rune(rand.Intn(26) + 65)
		chars[1] = rune(rand.Intn(26) + 65)
		chars[2] = rune(rand.Intn(10) + 48)
		chars[3] = rune(rand.Intn(10) + 48)
		chars[4] = rune(rand.Intn(10) + 48)
		n = string(chars)

		unique = !robotRegister[n]

	}

	robotRegister[n] = true
	return n
}
