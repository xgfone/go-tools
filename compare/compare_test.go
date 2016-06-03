package compare_test

import (
	"testing"

	"github.com/xgfone/go-tools/compare"
)

func TestCompare(t *testing.T) {
	v1 := []uint16{1, 2, 4}
	v2 := []uint16{1, 2, 3}
	if !compare.GT(v1, v2) {
		t.Fail()
	}

	if !compare.LT(v2, v1) {
		t.Fail()
	}

	if compare.EQ(v1, v2) {
		t.Fail()
	}

	if !compare.EQ([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}

	if compare.LT([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}
}
