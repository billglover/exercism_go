package scrabble

import "strings"

const testVersion = 5

var letterScores = map[int][]rune{
	1:  {'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	2:  {'D', 'G'},
	3:  {'B', 'C', 'M', 'P'},
	4:  {'F', 'H', 'V', 'W', 'Y'},
	5:  {'K'},
	8:  {'J', 'X'},
	10: {'Q', 'Z'},
}

var scoreMap = make(map[rune]int, 26)

func init() {

	// create a more efficient representation of our letter scores
	for s, letters := range letterScores {
		for _, l := range letters {
			scoreMap[l] = s
		}
	}
}

// Score reutrns the Scrabble score for a given word. Invalid
// characters and whitespace will be treated as scoring 0.
func Score(w string) (score int) {

	w = strings.ToUpper(w)

	for _, c := range w {
		i, ok := scoreMap[c]
		if ok {
			score += i
		}
	}

	return score
}
