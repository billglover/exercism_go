package strain

// Ints defines a collection of int
type Ints []int

// Lists defines a collection of Ints
type Lists [][]int

// Strings defines a collection of string
type Strings []string

// Keep returns a new collection containing elements where the predicate is true
func (i Ints) Keep(f func(int) bool) Ints {
	if i == nil {
		return i
	}

	r := make([]int, len(i))
	ri := 0

	for x := range i {
		if f(i[x]) {
			r[ri] = i[x]
			ri++
		}
	}

	return r[:ri]
}

// Discard returns a new collection containing elements where the predicate is false
func (i Ints) Discard(f func(int) bool) Ints {
	r := i.Keep(func(x int) bool { return f(x) != true })
	return r
}

// Keep returns a new collection containing elements where the predicate is true
func (l Lists) Keep(f func([]int) bool) Lists {
	if l == nil {
		return l
	}

	r := make([][]int, len(l))
	ri := 0

	for x := range l {
		if f(l[x]) {
			r[ri] = l[x]
			ri++
		}
	}

	return r[:ri]
}

// Keep returns a new collection containing elements where the predicate is true
func (s Strings) Keep(f func(string) bool) Strings {
	if s == nil {
		return s
	}

	r := make([]string, len(s))
	ri := 0

	for x := range s {
		if f(s[x]) {
			r[ri] = s[x]
			ri++
		}
	}

	return r[:ri]
}
