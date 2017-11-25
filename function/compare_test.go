package function

import (
	"testing"
)

func TestCompare(t *testing.T) {
	v1 := []uint16{1, 2, 4}
	v2 := []uint16{1, 2, 3}
	if !GT(v1, v2) {
		t.Fail()
	}

	if !LT(v2, v1) {
		t.Fail()
	}

	if EQ(v1, v2) {
		t.Fail()
	}

	if !EQ([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}

	if LT([]int{1, 2, 3}, []int{1, 2, 3}) {
		t.Fail()
	}
}
