package stringset

import (
	"reflect"
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
	for k := range s1 {
		if ok := s2[k]; ok == true {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	is := Set{}

	for k := range s1 {
		if ok := s2[k]; ok == true {
			is[k] = true
		}
	}

	return is
}

func Difference(s1, s2 Set) Set {
	ds := Set{}

	for k := range s1 {
		if ok := s2[k]; ok == false {
			ds[k] = true
		}
	}

	return ds
}

func Union(s1, s2 Set) Set {
	us := s2

	for k := range s1 {
		us[k] = true
	}

	return us
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
