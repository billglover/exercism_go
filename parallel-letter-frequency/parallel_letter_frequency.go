package letter

const testVersion = 1

// ConcurrentFrequency takes a slice of strings and calculates the frequency
// of runes within the strings.
func ConcurrentFrequency(strings []string) FreqMap {

	m := FreqMap{}
	results := make(chan FreqMap)

	for _, part := range strings {
		go func(s string) {
			results <- Frequency(s)
		}(part)
	}

	for r := 3; r > 0; r-- {
		f := <-results
		for l := range f {
			m[l] += f[l]
		}
	}
	return m
}
