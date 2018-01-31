package railfence

import (
	"bytes"
	"strings"
)

const (
	up   = 1
	down = -1
)

// Encode takes a message and a key and returns a version of the message
// encoded with the rail fence cipher.
func Encode(msg string, key int) string {

	// This version of the cypher strips spaces from the message.
	msg = strings.Replace(msg, " ", "", -1)

	lines := make([][]byte, key)
	curLine := 0
	dir := up

	for c := range msg {
		lines[curLine] = append(lines[curLine], msg[c])
		switch curLine {
		case 0:
			dir = up
		case key - 1:
			dir = down
		}
		curLine += dir
	}

	code := bytes.Join(lines, []byte(""))
	return string(code)
}

// Decode takes an encoded message and a key and uses the rail fence
// cipher to return a decoded version of the message.
func Decode(msg string, key int) string {

	// calculate the length of each line
	lens := make([]int, key)
	curLine := 0
	dir := up

	for _ = range msg {
		lens[curLine]++
		switch curLine {
		case 0:
			dir = up
		case key - 1:
			dir = down
		}
		curLine += dir
	}

	// create slices representing each line
	lines := make([][]byte, key)
	ptr := 0
	for l := range lines {
		lines[l] = []byte((msg[ptr : ptr+lens[l]]))
		ptr += lens[l]
	}

	// Read the characters off the line
	pos := make([]int, key)
	curLine = 0
	dir = up
	plain := make([]byte, len(msg))

	for i := range plain {

		plain[i] = lines[curLine][pos[curLine]]
		pos[curLine]++

		switch curLine {
		case 0:
			dir = up
		case key - 1:
			dir = down
		}
		curLine += dir
	}

	return string(plain)
}
