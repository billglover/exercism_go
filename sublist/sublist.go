package sublist

type Relation string

func Sublist(a, b []int) Relation {

	// empty lists are regard as equal
	if len(a) == 0 && len(b) == 0 {
		return "equal"
	}

	// empty lists are a sublist of any non-empty list
	if len(a) == 0 && len(b) > 0 {
		return "sublist"
	}

	// lists that contain empty lists are superlists
	if len(a) > 0 && len(b) == 0 {
		return "superlist"
	}

	// look for identical lists
	if len(a) == len(b) {
		if isEqual(a, b) {
			return "equal"
		}
		return "unequal"
	}

	// look for sublists
	if len(b) > len(a) {
		mo := len(b) - len(a) + 1

		for o := 0; o < mo; o++ {
			if isEqual(a, b[o:(o+len(a))]) {
				return "sublist"
			}
		}
	}

	// look for superlists
	if len(a) > len(b) {
		mo := len(a) - len(b) + 1

		for o := 0; o < mo; o++ {
			if isEqual(b, a[o:(o+len(b))]) {
				return "superlist"
			}
		}
	}

	return "unequal"
}

func isEqual(a, b []int) bool {
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	return false
}
