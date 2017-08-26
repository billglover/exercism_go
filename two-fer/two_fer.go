package twofer

import "fmt"

// ShareWith takes a name and returns how things
// should be shared.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
