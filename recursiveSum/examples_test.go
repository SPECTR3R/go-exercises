package recursiveSum_test

import (
	"fmt"

	"github.com/SPECTR3R/go-exercises/recursiveSum"
)

func ExampleInts() {
	s := recursiveSum.Ints(1, 2, 3, 4, 5)
	fmt.Println("Sum of 1 to 5 is,", s)
	// Output:
	// Sum of 1 to 5 is, 15
}
