package main

import "fmt"

// START2 OMIT
func main() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	fmt.Println(first(numbers)) // HL

	names := []string{"foo", "bar", "baz", "qux", "quux"}
	fmt.Println(first(names)) // HL

	facts := []bool{false, true, false, false}
	fmt.Println(first(facts)) // HL
}
// END2 OMIT

// START1 OMIT
func first[T any](s []T) (T, bool) { // HL
	if len(s) == 0 {
		var zero T // HL
		return zero, false
	}
	return s[0], true
}

// END1 OMIT
