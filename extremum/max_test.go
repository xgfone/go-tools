package extremum_test

import (
	"testing"

	"github.com/xgfone/go-tools/extremum"
)

func TestMax(t *testing.T) {
	if extremum.Max(11, 22) != 22 {
		t.Fail()
	}

	v := []int64{1, 2, 3, 4}
	if extremum.MaxInSlice(v) != 3 {
		t.Fail()
	}
}
