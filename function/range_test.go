package function

import (
	"fmt"
)

func ExampleRange() {
	fmt.Println(Range(1, 10, 2))
	fmt.Println(Range(10, 1, -2))

	// Output:
	// [1 3 5 7 9]
	// [10 8 6 4 2]
}

func ExampleRangeWithStep() {
	fmt.Println(RangeWithStep(2)(1, 10))
	fmt.Println(RangeWithStep(-2)(10, 1))

	// Output:
	// [1 3 5 7 9]
	// [10 8 6 4 2]
}

func ExampleRangeByStop() {
	fmt.Println(RangeByStop(1, 2)(10))
	fmt.Println(RangeByStop(10, -2)(1))

	// Output:
	// [1 3 5 7 9]
	// [10 8 6 4 2]
}
