package scale

import "strings"

var (
	sharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	flat  = []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"}
	flats = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
)

// Scale takes a tonic and a set of intervals and generates a musical scale.
func Scale(tonic, interval string) []string {

	// Pick a representation of the chromatic scale. By default we use sharps
	// but certain tonics require the use of the flat representation.
	base := sharp
	for i := range flats {
		if flats[i] == tonic {
			base = flat
		}
	}
	baseLen := len(base)

	// Identify the start point in the chromatic scale.
	tonic = strings.Title(tonic)
	index, found := 0, false
	for index = range base {
		if base[index] == tonic {
			found = true
			break
		}
	}
	if found == false {
		return nil
	}

	// If the interval is not defined, the default is to use every note in the scale.
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}
	scale := make([]string, len(interval))

	// The initial note in the scale is always the start point in the base scale.
	offset := 0
	scale[0] = base[(index)]

	for i := 0; i < len(interval)-1; i++ {
		switch interval[i] {
		case 'M':
			offset = offset + 2
		case 'm':
			offset = offset + 1
		case 'A':
			offset = offset + 3
		}
		scale[i+1] = base[(index+offset)%baseLen]
	}

	return scale
}
