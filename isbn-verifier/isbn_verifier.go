package isbn

// IsValidISBN takes a string and determines whether it is a valid ISBN10 number.
func IsValidISBN(s string) bool {

	sum := 0
	count := 0

	for i := 0; i < len(s); i++ {
		switch {

		case count > 9:
			return false

		case s[i] >= '0' && s[i] <= '9':
			sum = sum + (int(s[i]-48) * (10 - count))
			count++

		case s[i] == '-':
			continue

		case s[i] == 'X':
			if count != 9 {
				return false
			}
			sum = sum + 10
			count++

		default:
			return false
		}
	}

	if count <= 9 {
		return false
	}

	valid := sum%11 == 0
	return valid
}
