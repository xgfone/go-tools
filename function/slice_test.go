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

func ExampleReverse() {
	// []interface{}
	s1 := []interface{}{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(s1))
	s1 = []interface{}{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Reverse(s1))

	// []string
	s2 := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(Reverse(s2))
	s2 = []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(Reverse(s2))

	// []int
	s3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(s3))
	s3 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Reverse(s3))

	// []int64
	s4 := []int64{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(s4))
	s4 = []int64{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Reverse(s4))

	// []uint
	s5 := []uint{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(s5))
	s5 = []uint{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Reverse(s5))

	// []uint64
	s6 := []uint64{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(s6))
	s6 = []uint64{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Reverse(s6))

	// Output:
	// [6 5 4 3 2 1]
	// [7 6 5 4 3 2 1]
	// [f e d c b a]
	// [g f e d c b a]
	// [6 5 4 3 2 1]
	// [7 6 5 4 3 2 1]
	// [6 5 4 3 2 1]
	// [7 6 5 4 3 2 1]
	// [6 5 4 3 2 1]
	// [7 6 5 4 3 2 1]
	// [6 5 4 3 2 1]
	// [7 6 5 4 3 2 1]
}
