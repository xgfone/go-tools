package extremum

import (
	"testing"
)

func TestMin(t *testing.T) {
	if Min(11, 22).(int) != 11 {
		t.Fail()
	}

	v := []int64{1, 2, 3, 4}
	if MinInSlice(v).(int64) != 1 {
		t.Fail()
	}
}
