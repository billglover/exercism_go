package stringset

import (
	"strings"
)

const (
	testVersion = 4
)

type Set map[string]bool

func New() Set {
	ss := Set{}
	return ss
}

func NewFromSlice(sl []string) Set {
	ss := &Set{}
	for _, s := range sl {
		ss.Add(s)
	}
	return *ss
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if ok := s2[k]; ok == false {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return false
}

func Equal(s1, s2 Set) bool {
	return false
}

func Intersection(s1, s2 Set) Set {
	return nil
}

func Difference(s1, s2 Set) Set {
	return nil
}

func Union(s1, s2 Set) Set {
	return nil
}

func (ss Set) String() string {
	s := "{"

	vals := make([]string, len(ss))
	i := 0
	for v := range ss {
		vals[i] = "\"" + v + "\""
		i++
	}

	s += strings.Join(vals, ", ")

	s += "}"
	return s
}

func (ss *Set) Has(v string) bool {
	_, has := (*ss)[v]
	return has
}

func (ss *Set) IsEmpty() bool {
	return len(*ss) == 0
}

func (ss *Set) Add(s string) {
	(*ss)[s] = true
}
