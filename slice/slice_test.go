package slice_test

import (
	"fmt"
	"testing"

	"github.com/xgfone/go-tools/slice"
)

func TestSlice(t *testing.T) {
	ss := []int{1, 2, 3, 4, 5, 6}
	out := 0

	ok := slice.SetValueWithDefault(&out, ss, 2, 100)
	if ok != nil || out != 3 {
		t.Fail()
	}

	ok = slice.SetValueWithDefault(&out, ss, 6, 100)
	if ok != nil || out != 100 {
		t.Fail()
	}

	ok = slice.SetValue(out, ss, 6)
	if ok == nil {
		t.Fail()
	}
}

func ExampleSetValue() {
	ss := []int{1, 2, 3, 4, 5, 6}
	out := -1
	err := slice.SetValue(&out, ss, 1)
	fmt.Println(out, err)

	out = -1
	err = slice.SetValue(&out, ss, 6)
	fmt.Println(out, err)
	// Output:
	// 2 <nil>
	// -1 The invalid index
}

func ExampleSetValueWithDefault() {
	ss := []int{1, 2, 3, 4, 5, 6}
	out := 0
	err := slice.SetValueWithDefault(&out, ss, 1, -1)
	fmt.Println(out, err)

	out = 0
	err = slice.SetValueWithDefault(&out, ss, 6, -1)
	fmt.Println(out, err)
	// Output:
	// 2 <nil>
	// -1 <nil>
}

func ExampleIn() {
	ss := []string{"hello", "world"}
	fmt.Println(slice.In("hello", ss))
	fmt.Println(slice.In("aaron", ss))

	// Output:
	// true
	// false
}
