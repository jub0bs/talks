package main

import (
	"fmt"
)


// START OMIT
type Plusser interface {
  ~int | ~uint | ~string
}

func main() {
	var p Plusser // compilation error // HL
	fmt.Println(p)
}

// END OMIT
