package extremum_test

import (
	"testing"

	"github.com/xgfone/go-tools/extremum"
)

func TestMin(t *testing.T) {
	if extremum.Min(11, 22) != 11 {
		t.Fail()
	}

	v := []int64{1, 2, 3, 4}
	if extremum.MinInSlice(v) != 0 {
		t.Fail()
	}
}
