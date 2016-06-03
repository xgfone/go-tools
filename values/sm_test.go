package values_test

import (
	"testing"

	"github.com/xgfone/go-tools/values"
)

func TestSlice(t *testing.T) {
	ms := []interface{}{"test"}
	ss := []interface{}{11, "aa", values.SMap{"abcd": 11, "ok": false, "slice": values.Slice(ms)}, true}
	s := values.Slice(ss)

	if _, ok := s.Int(0); !ok {
		t.Fail()
	}

	if _, ok := s.String(1); !ok {
		t.Fail()
	}

	if v, ok := s.SMap(2); !ok {
		t.Fail()
	} else {
		if _, ok := v.Int("abcd"); !ok {
			t.Fail()
		}

		if _, ok := v.Bool("ok"); !ok {
			t.Fail()
		}

		if _s, ok := v.Slice("slice"); !ok {
			t.Fail()
		} else {
			if _, ok := _s.String(0); !ok {
				t.Fail()
			}
		}
	}

	if _, ok := s.Bool(3); !ok {
		t.Fail()
	}
}
