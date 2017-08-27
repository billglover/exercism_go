package summultiples

const testVersion = 2

// SumMultiples returns the sum of all multiples of a number
// up to, but not including, the limit specified.
func SumMultiples(limit int, multiples ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, m := range multiples {
			if i%m == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
