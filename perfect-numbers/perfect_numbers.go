package perfect

import (
	"errors"
)

const testVersion = 1

// Classification represents the classification of a number
// based on the aliquot sum of that number.
type Classification int

// Constants
const (
	ClassificationUnknown   Classification = iota // unknown classification
	ClassificationDeficient                       // the aliquot sum < the number
	ClassificationAbundant                        // the aliquot sum > the number
	ClassificationPerfect                         // the aliquot sum = the number
)

// String returns a friendly representation of the
// classification useful for debugging.
func (c Classification) String() string {
	switch c {
	case ClassificationDeficient:
		return "deficient"
	case ClassificationAbundant:
		return "abundant"
	case ClassificationPerfect:
		return "perfect"
	default:
		return "unknown"
	}
}

// ErrOnlyPositive is an error that indicates a non positive integer has
// been passed.
var ErrOnlyPositive = errors.New("only positive integers accepted")

// Classify takes in an unsigned integer and classifies it according to
// the aliquot sum of that number. It returns an error if a non positive
// integer was passed to the function.
func Classify(n uint64) (c Classification, err error) {
	if n <= 0 {
		return ClassificationUnknown, ErrOnlyPositive
	}

	sof := sumOfFactors(n)
	switch {
	case sof > n:
		return ClassificationAbundant, nil
	case sof < n:
		return ClassificationDeficient, nil
	case sof == n:
		return ClassificationPerfect, nil
	}
	return ClassificationPerfect, nil
}

// sumOfFactors takes an unsigned integer and reutrns the sum of
// all the factors up to but not including the number itself.
func sumOfFactors(n uint64) (sum uint64) {
	for f := uint64(1); f < n; f++ {
		if n%f == 0 {
			sum += f
		}
	}
	return sum
}
