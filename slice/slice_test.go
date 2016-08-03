package slice_test

import (
	"fmt"
	"testing"

	"github.com/xgfone/go-tools/slice"
)

func TestSlice(t *testing.T) {
	ss := []int{1, 2, 3, 4, 5, 6}
	ok := true
	out := 0

	ok = slice.SetValueWithDefault(&out, ss, 2, 100)
	if !ok || out != 3 {
		t.Fail()
	}

	ok = slice.SetValueWithDefault(&out, ss, 6, 100)
	if !ok || out != 100 {
		t.Fail()
	}

	ok = slice.SetValue(out, ss, 6)
	if ok {
		t.Fail()
	}
}

func ExampleSetValue() {
	ss := []int{1, 2, 3, 4, 5, 6}
	out := -1
	ok1 := slice.SetValue(&out, ss, 1)
	fmt.Println(out, ok1)

	out = -1
	ok2 := slice.SetValue(&out, ss, 6)
	fmt.Println(out, ok2)
	// Output:
	// 2 true
	// -1 false
}

func ExampleSetValueWithDefault() {
	ss := []int{1, 2, 3, 4, 5, 6}
	out := 0
	ok1 := slice.SetValueWithDefault(&out, ss, 1, -1)
	fmt.Println(out, ok1)

	out = 0
	ok2 := slice.SetValueWithDefault(&out, ss, 6, -1)
	fmt.Println(out, ok2)
	// Output:
	// 2 true
	// -1 true
}

func ExampleIn() {
	ss := []string{"hello", "world"}
	fmt.Println(slice.In("hello", ss))
	fmt.Println(slice.In("aaron", ss))

	// Output:
	// true
	// false
}
