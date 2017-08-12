package accumulate

const testVersion = 1

// Accumulate takes a string slice and a conversion function
// and returns a new string slice containing the result of
// applying that conversion to each element of the input.
func Accumulate(c []string, f func(string) string) []string {
	r := make([]string, len(c))
	for i, v := range c {
		r[i] = f(v)
	}
	return r
}
