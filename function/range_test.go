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

func ExampleRanges() {
	fmt.Println(Ranges(10))
	fmt.Println(Ranges(1, 10))
	fmt.Println(Ranges(1, 10, 1))
	fmt.Println(Ranges(1, 10, 2))
	fmt.Println(Ranges(10, 0))

	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
	// [1 2 3 4 5 6 7 8 9]
	// [1 3 5 7 9]
	// []
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

func ExampleRangeStepOne() {
	fmt.Println(RangeStepOne(1, 10))

	// Output:
	// [1 2 3 4 5 6 7 8 9]
}

func ExampleRangeStop() {
	fmt.Println(RangeStop(10))

	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}
