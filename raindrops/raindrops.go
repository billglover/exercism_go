package raindrops

import "fmt"

const testVersion = 3

// Convert takes a number and converts it to a string
// by applying the PlingPlangPlong conversion.
// - If the number has 3 as a factor, output 'Pling'.
// - If the number has 5 as a factor, output 'Plang'.
// - If the number has 7 as a factor, output 'Plong'.
// - If the number does not have 3, 5, or 7 as a factor,
//   just pass the number's digits straight through.
func Convert(i int) (result string) {

	if i%3 == 0 {
		result += "Pling"
	}

	if i%5 == 0 {
		result += "Plang"
	}

	if i%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = fmt.Sprintf("%d", i)
	}

	return result
}
