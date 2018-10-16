package ocr

import (
	"strings"
)

// Recognize takes a string representing a 3 x 4 grid of pipes, underscores,
// and spaces, determines what series of numbers it represents. It returns
// a slice of strings containing the numbers found. A ? is used in place of
// any unrecognizeable numbers.
func Recognize(s string) []string {

	lines := strings.Split(s, "\n")
	text := []string{}

	for l := 1; l < len(lines); l = l + 4 {
		number := ""
		for c := 0; c < len(lines[l]); c = c + 3 {
			number = number + recognizeDigit(lines[l][c:c+3]+
				lines[l+1][c:c+3]+
				lines[l+2][c:c+3]+
				lines[l+3][c:c+3])
		}
		text = append(text, number)
	}

	return text
}

func recognizeDigit(s string) string {
	s = strings.Replace(s, "\n", "", 4)
	d, ok := pattern[s]
	if ok == false {
		return "?"
	}
	return d
}

var pattern = map[string]string{
	" _ | ||_|   ": "0",
	"     |  |   ": "1",
	" _  _||_    ": "2",
	" _  _| _|   ": "3",
	"   |_|  |   ": "4",
	" _ |_  _|   ": "5",
	" _ |_ |_|   ": "6",
	" _   |  |   ": "7",
	" _ |_||_|   ": "8",
	" _ |_| _|   ": "9",
}
