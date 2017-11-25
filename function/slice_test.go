package function

import (
	"fmt"
)

func ExamplePullSliceValue() {
	ss := []int{1, 2, 3, 4, 5, 6}
	out := -1
	err := PullSliceValue(&out, ss, 1)
	fmt.Println(out, err)

	out = -1
	err = PullSliceValue(&out, ss, 6)
	fmt.Println(out, err)

	// Output:
	// 2 <nil>
	// -1 the index is exceeds the length of the slice
}

func ExamplePullSliceValueWithDefault() {
	ss := []int{1, 2, 3, 4, 5, 6}
	out := 0
	err := PullSliceValueWithDefault(&out, ss, 1, -1)
	fmt.Println(out, err)

	out = 0
	err = PullSliceValueWithDefault(&out, ss, 6, -1)
	fmt.Println(out, err)

	// Output:
	// 2 <nil>
	// -1 <nil>
}

func ExampleInSlice() {
	ss := []interface{}{"hello", "world"}
	fmt.Println(InSlice("hello", ss))
	fmt.Println(InSlice("aaron", ss))

	// Output:
	// true
	// false
}
