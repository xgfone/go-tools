package values_test

import (
	"sort"
	"testing"

	"github.com/xgfone/go-tools/values"
)

func TestSlice(t *testing.T) {
	ms := []interface{}{"test"}
	ss := []interface{}{11, "aa", values.SMap{"abcd": 11, "ok": false, "slice": values.Slice(ms)}, true}
	s := values.Slice(ss)

	if v, err := s.Int(0); err != nil || v != 11 {
		t.Fail()
	}

	if v, err := s.String(1); err != nil || v != "aa" {
		t.Fail()
	}

	if v, err := s.SMap(2); err != nil {
		t.Fail()
	} else {
		if vv, err := v.Int("abcd"); err != nil || vv != 11 {
			t.Fail()
		}

		if vv, err := v.Bool("ok"); err != nil || vv != false {
			t.Fail()
		}

		if _s, err := v.Slice("slice"); err != nil {
			t.Fail()
		} else {
			if v, err := _s.String(0); err != nil || v != "test" {
				t.Fail()
			}
		}
	}

	if _, err := s.Bool(3); err != nil {
		t.Fail()
	}

	if sm, err := s.SMap(2); err != nil {
		t.Fail()
	} else {
		ks := sm.Keys()
		if sort.StringSlice(ks).Sort(); ks[2] != "slice" {
			t.Fail()
		}
	}

	mi := map[string]int{"aa": 11, "bb": 22}
	if sm, err := values.ConvertToSMap(mi); err != nil || sm["aa"].(int) != 11 || sm["bb"].(int) != 22 {
		t.Fail()
	}

	si := []int{11, 22}
	if ss, err := values.ConvertToSlice(si); err != nil || ss[0].(int) != 11 || ss[1].(int) != 22 {
		t.Fail()
	}

	if _, err := values.ConvertToSMap(nil); err == nil {
		t.Fail()
	}

	if _, err := values.ConvertToSlice(nil); err == nil {
		t.Fail()
	}
}
