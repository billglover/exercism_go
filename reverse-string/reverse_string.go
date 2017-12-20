// Package reverse provides function for reversing
// the contents of various data structures.
package reverse

// String takes a string and returns the reverse
func String(s string) string {
	b := []byte(s)

	// Loop from both ends of the byte slice at once.
	// This may not be as readable as declaring a single
	// loop, but it is compact.
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
