package function

import (
	"fmt"
)

func ExampleRange() {
	fmt.Println(Ranges(1, 10, 2))
	fmt.Println(Ranges(10, 1, -2))

	// Output:
	// [1 3 5 7 9]
	// [10 8 6 4 2]
}

func ExampleRanges() {
	fmt.Println(Range(10))
	fmt.Println(Range(1, 10))
	fmt.Println(Range(1, 10, 1))
	fmt.Println(Range(1, 10, 2))
	fmt.Println(Range(10, 0))

	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
	// [1 3 5 7 9]
	// []
}
