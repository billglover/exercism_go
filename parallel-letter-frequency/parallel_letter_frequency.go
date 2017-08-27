package letter

const testVersion = 1

// ConcurrentFrequency takes a slice of strings and calculates the frequency
// of runes within the strings.
func ConcurrentFrequency(strings []string) FreqMap {

	m := FreqMap{}
	results := make(chan FreqMap, len(strings))

	for _, part := range strings {
		go func(s string) {
			results <- Frequency(s)
		}(part)
	}

	for r := len(strings); r > 0; r-- {
		f := <-results
		for l := range f {
			m[l] += f[l]
		}
	}
	return m
}
