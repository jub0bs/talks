package main

import (
	"fmt"
)

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


// START OMIT
func (s Set[T]) Map[U any](func(T) U) Set[U] { // compilation error // HL
	res := make(Set[U], len(s))
	for t := range s {
		res.Add(f(t))
	}
	return res
}

func main() {
	stringSet := NewSet[string]()
	stringSet.Add("foo")
	stringSet.Add("bar")
	f := func(s string) int { return len(s) }
	intSet := stringSet.Map(f) // HL
	fmt.Println(intSet)
}

// END OMIT
