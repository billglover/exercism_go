package strand

const testVersion = 3

var nucleotides = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA takes a DNA strand and returns its RNA complement
func ToRNA(dna string) (rna string) {
	r := make([]rune, len(dna))
	for i, n := range dna {
		r[i] = nucleotides[n]
	}
	return string(r)
}
