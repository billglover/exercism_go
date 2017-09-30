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
	for _, n := range dna {
		rna += string(nucleotides[n])
	}
	return rna
}
