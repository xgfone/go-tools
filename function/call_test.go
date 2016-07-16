package function_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/xgfone/go-tools/function"
)

func get(i int, j int) (int, error) {
	return i + j, nil
}

func TestCall(t *testing.T) {
	if ret, err := function.Call(get, 1, 2); err != nil {
		t.Fail()
	} else {
		if ret[0].(int) != 3 || ret[1] != nil {
			t.Fail()
		}
	}
}

func TestCallWithPointer(t *testing.T) {
	f := func(v *int) (old int) {
		old = *v
		*v += 1
		return
	}

	v := 1
	ret, _ := function.Call(f, &v)
	// The returned value is the old, which is 1, and v became 2.
	if ret[0].(int) != 1 || v != 2 {
		t.Fail()
	}
}

func ExampleCall() {
	f := func(i int, j int) (int, error) {
		return i + j, errors.New("This is not an error")
	}

	ret, _ := function.Call(f, 1, 2)

	// Since the first result is an integer, and it's not necessary to check
	// whether it is nil, so you may omit it, and infer this type directly.
	if ret[0] != nil {
		fmt.Println(ret[0].(int))
	}

	// Since the second result may be nil, so you MUST check whether it is nil firstly.
	if ret[1] != nil {
		fmt.Println(ret[1].(error))
	}
	// Output:
	// 3
	// This is not an error
}
