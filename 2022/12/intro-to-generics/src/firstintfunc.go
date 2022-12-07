package main

import "fmt"

// START OMIT
func main() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	num, ok := first(numbers) // HL
	if !ok {
		fmt.Println("empty slice")
		return
	}
	fmt.Println("first element:", num)
}

func first(s []int) (int, bool) { // HL
	if len(s) == 0 {
		return 0, false
	}
	return s[0], true
}

// END OMIT
