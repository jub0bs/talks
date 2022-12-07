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
type Plusser interface {
  ~int | ~uint | ~string
}

func Sum[E Plusser](s Set[E]) E { // HL
	var res E
	for e := range s {
		res += e
	}
	return res
}

func main() {
	set := NewSet[int]()
	set.Add(4)
	set.Add(666)
	fmt.Println(Sum(set)) // HL
}

// END OMIT
