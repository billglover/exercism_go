//Package letter calculaes the letter frequency in a given set of strings
package letter

import "sync"

const testVersion = 1

// ConcurrentFrequency takes a slice of strings and calculates the frequency
// of runes within the strings.
func ConcurrentFrequency(strings []string) FreqMap {
	fMap := FreqMap{}
	rChan := make(chan FreqMap, len(strings))

	var wg sync.WaitGroup
	wg.Add(len(strings))

	for _, item := range strings {
		go func(s string) {
			defer wg.Done()
			rChan <- Frequency(s)
		}(item)
	}

	go func() {
		wg.Wait()
		close(rChan)
	}()

	for f := range rChan {
		for l := range f {
			fMap[l] += f[l]
		}
	}

	return fMap
}
