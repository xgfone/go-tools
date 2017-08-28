package extremum

import "testing"

func TestMax(t *testing.T) {
	if Max(11, 22).(int) != 22 {
		t.Fail()
	}

	v := []int64{1, 2, 3, 4}
	if MaxInSlice(v).(int64) != 4 {
		t.Fail()
	}
}
