package main

import "fmt"

func main() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	names := []string{"foo", "bar", "baz", "qux", "quux"}
	fmt.Println(first(numbers))
	fmt.Println(first(names))
}

// START OMIT
func first(s interface{}) (interface{}, bool) { // HL
	switch s := s.(type) {
	case []int:
		if len(s) == 0 {
			return 0, false
		}
		return s[0], true
	// cases for other slice types omitted
	default:
		// what to do in this case is unclear...
		panic(fmt.Sprintf("unsupported type: %T", s))
	}
}

// END OMIT
