package main

import (
	"fmt"
)

// START1 OMIT
type Set[E comparable] map[E]struct{} // HL

func NewSet[E comparable]() Set[E] { // HL
	return make(map[E]struct{})
}

func (s Set[E]) Contains(e E) bool { // HL
	_, found := s[e]
	return found
}

func (s Set[E]) Add(e E) { // HL
	s[e] = struct{}{}
}

func (s Set[E]) Remove(e E) { // HL
	delete(s, e)
}

// END1 OMIT

// START2 OMIT
func main() {
	intSet := NewSet[int]() // HL
	intSet.Add(4)
	intSet.Add(666)
	fmt.Println(intSet)

	stringSet := NewSet[string]() // HL
	stringSet.Add("foo")
	stringSet.Add("bar")
	fmt.Println(stringSet)
}

// END2 OMIT
