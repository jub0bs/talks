package main

import "fmt"

// START OMIT
func main() {
	names := []string{"foo", "bar", "baz", "qux", "quux"}
	name, ok := firstStringElem(names) // HL
	if !ok {
		fmt.Println("empty slice")
		return
	}
	fmt.Println("first element:", name)
}

func firstStringElem(s []string) (string, bool) { // HL
	if len(s) == 0 {
		return "", false
	}
	return s[0], true
}

// END OMIT
