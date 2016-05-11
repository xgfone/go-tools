package slice_test

import (
    "testing"

    "github.com/xgfone/go-tools/slice"
)

func TestSlice(t *testing.T) {
    ss := []int{1,2,3,4,5,6}
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