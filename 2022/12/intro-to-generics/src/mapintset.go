package main

import (
	"fmt"
)

// START1 OMIT
type Set map[int]struct{} // HL

func NewSet() Set { // HL
	return make(map[int]struct{})
}

func (s Set) Contains(e int) bool { // HL
	_, found := s[e]
	return found
}

func (s Set) Add(e int) { // HL
	s[e] = struct{}{}
}

func (s Set) Remove(e int) { // HL
	delete(s, e)
}

// END1 OMIT

// START2 OMIT
func main() {
	set := NewSet() // HL
	fmt.Println(set)

	k := 666
	if !set.Contains(k) { // HL
		fmt.Printf("%d is not in the map.\n", k)
	}

	set.Add(4) // HL
	set.Add(666)
	fmt.Println(set)

	set.Remove(4) // HL
	fmt.Println(set)
}

// END2 OMIT
