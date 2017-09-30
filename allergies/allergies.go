package allergies

const testVersion = 1

// AllergyScores maps each allergy onto its
// appropriate numberic score.
var AllergyScores = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns the list of alllergies based
// on a single numeric allergy score.
func Allergies(s uint) (allergies []string) {
	for name := range AllergyScores {
		if AllergicTo(s, name) {
			allergies = append(allergies, name)
		}
	}

	return allergies
}

// AllergicTo determines whether someone is allergic
// to a specific allergy based on an allergy score.
func AllergicTo(s uint, a string) bool {
	return (s&AllergyScores[a] > 0)
}
