package values_test

import (
	"testing"

	"github.com/xgfone/go-tools/values"
)

func TestSlice(t *testing.T) {
	ms := []interface{}{"test"}
	ss := []interface{}{11, "aa", values.SMap{"abcd": 11, "ok": false, "slice": values.Slice(ms)}, true}
	s := values.Slice(ss)

	if v, ok := s.Int(0); !ok || v != 11 {
		t.Fail()
	}

	if v, ok := s.String(1); !ok || v != "aa" {
		t.Fail()
	}

	if v, ok := s.SMap(2); !ok {
		t.Fail()
	} else {
		if vv, ok := v.Int("abcd"); !ok || vv != 11 {
			t.Fail()
		}

		if vv, ok := v.Bool("ok"); !ok || vv != false {
			t.Fail()
		}

		if _s, ok := v.Slice("slice"); !ok {
			t.Fail()
		} else {
			if v, ok := _s.String(0); !ok || v != "test" {
				t.Fail()
			}
		}
	}

	if _, ok := s.Bool(3); !ok {
		t.Fail()
	}
}
