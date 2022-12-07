package main

import "fmt"

// START OMIT
func main() {
	numbers := []int{4, 8, 15, 16, 23, 42}
	if len(numbers) == 0 {
		fmt.Println("empty slice")
		return
	}
	fmt.Println("first element:", numbers[0])
}

// END OMIT
