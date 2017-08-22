package series

const testVersion = 2

// All returns all the contiguous substrings of length `n` from
// a given a string of digits.
func All(n int, s string) (series []string) {
	for i := 0; i < len(s) && i+n <= len(s); i++ {
		ns, ok := First(n, s[i:])
		if ok == false {
			return nil
		}
		series = append(series, ns)
	}

	return series
}

// UnsafeFirst returns the first contiguous substrings of length `n` from
// a given a string of digits. It does no bounds checking and will crash
// if `n` is greater than the length of the string.
func UnsafeFirst(n int, s string) (f string) {
	return s[:n]
}

// First returns the first contiguous substrings of length `n` from
// a given a string of digits.
func First(n int, s string) (f string, ok bool) {
	if n > len(s) {
		return f, false
	}
	return s[:n], true
}
