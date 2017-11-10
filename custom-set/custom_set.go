// Package stringset implements a custom set type. The stringset
// takes string elements and provides typical set operations.
package stringset

import (
	"reflect"
	"strings"
)

const (
	testVersion = 4
)

// Set is a set of unique elements. Each element is a string
type Set map[string]bool

// New returns a new (empty) set.
func New() Set {
	ss := Set{}
	return ss
}

// NewFromSlice takes a slice of strings and returns a new set
// with each item in the slice representing one item in the set.
func NewFromSlice(sl []string) Set {
	ss := &Set{}
	for _, s := range sl {
		ss.Add(s)
	}
	return *ss
}

// Subset returns true if all elements in s1 are also present in s2.
func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if ok := s2[k]; ok == false {
			return false
		}
	}
	return true
}

// Disjoint returns true if both sets share no elements.
func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if ok := s2[k]; ok == true {
			return false
		}
	}
	return true
}

// Equal returns true if both sets are equal.
func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

// Intersection takes two Sets and returns a new Set containing the
// elements that are present in both.
func Intersection(s1, s2 Set) Set {
	is := Set{}

	for k := range s1 {
		if ok := s2[k]; ok == true {
			is[k] = true
		}
	}

	return is
}

// Difference takes two Sets and returns values from s1 that are
// not present in s2.
func Difference(s1, s2 Set) Set {
	ds := Set{}

	for k := range s1 {
		if ok := s2[k]; ok == false {
			ds[k] = true
		}
	}

	return ds
}

// Union takes two Sets and returns a new Set containing all
// elements from both sets.
func Union(s1, s2 Set) Set {
	us := s2

	for k := range s1 {
		us[k] = true
	}

	return us
}

// String returns a string representation of a Set.
func (ss Set) String() string {
	s := "{"

	el := make([]string, len(ss))
	i := 0
	for v := range ss {
		el[i] = "\"" + v + "\""
		i++
	}

	s += strings.Join(el, ", ")

	s += "}"
	return s
}

// Has returns true if a Set contains the provided element.
func (ss *Set) Has(v string) bool {
	_, has := (*ss)[v]
	return has
}

// IsEmpty returns true if a set contains no elements.
func (ss *Set) IsEmpty() bool {
	return len(*ss) == 0
}

// Add takes a string and adds an element to the set.
func (ss *Set) Add(s string) {
	(*ss)[s] = true
}
