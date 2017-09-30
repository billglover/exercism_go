package strand

const testVersion = 3

// ToRNA takes a DNA strand and returns its RNA complement
func ToRNA(dna string) (rna string) {
	r := make([]rune, len(dna))
	for i, n := range dna {
		switch n {
		case 'G':
			r[i] = 'C'
		case 'C':
			r[i] = 'G'
		case 'T':
			r[i] = 'A'
		case 'A':
			r[i] = 'U'
		}
	}
	return string(r)
}
