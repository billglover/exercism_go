// Package listops provides a number of basic list operations
package listops

// IntList is a list of integers
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Length returns the length of a list
func (l IntList) Length() int { return len(l) }

// Foldl takes a binary function and an initital value and returns the result
// of the fold left operation.
func (l IntList) Foldl(fn binFunc, initial int) int {
	result := initial
	for i := 0; i < len(l); i++ {
		result = fn(result, l[i])
	}
	return result
}

// Foldr takes a binary function and an initital value and returns the result
// of the fold right operation.
func (l IntList) Foldr(fn binFunc, initial int) int {
	result := initial
	for i := len(l) - 1; i >= 0; i-- {
		result = fn(l[i], result)
	}
	return result
}

// Filter takes a filter function and returns a new filtered list.
func (l IntList) Filter(fn predFunc) IntList {
	result := make(IntList, len(l))
	n := 0
	for i := range l {
		if fn(l[i]) == true {
			result[n] = l[i]
			n++
		}
	}
	return result[:n]
}

// Map takes a map function and returns a new list formed by applying the
// map function to each element in the original list.
func (l IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, len(l))
	for i := range l {
		result[i] = fn(l[i])
	}
	return result
}

// Reverse returns a new list formed by reversing the elements in the
// original list.
func (l IntList) Reverse() IntList {
	result := make(IntList, len(l))
	for i := range l {
		result[i] = l[len(l)-i-1]
	}
	return result
}

// Append takes a list and appends it to the original. The result is returned
// as a new list.
func (l IntList) Append(a IntList) IntList {
	result := make(IntList, len(l)+len(a))
	copy(result, l)
	for i := 0; i < len(a); i++ {
		result[len(l)+i] = a[i]
	}
	return result
}

// Concat takes a slice of lists and appends then to the original. The result
// is returned as a new list.
func (l IntList) Concat(a []IntList) IntList {
	result := make(IntList, len(l))
	copy(result, l)
	for i := range a {
		result = result.Append(a[i])
	}
	return result
}
