package etl

import "strings"

const testVersion = 1

// Transform takes a score indexed representation of letterScores
// and returns a letter indexed representation of the scores.
func Transform(letterScores map[int][]string) map[string]int {

	transform := make(map[string]int, 26)

	for s, letters := range letterScores {
		for _, l := range letters {
			l = strings.ToLower(l)
			transform[l] = s
		}
	}

	return transform
}
