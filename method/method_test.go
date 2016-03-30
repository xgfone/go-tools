package method_test

import (
	"testing"

	"github.com/xgfone/go-tools/method"
)

type T struct {
	Id   int
	Name string
}

func (t T) Get() (int, string) {
	return t.Id, t.Name
}

func TestMethod(t *testing.T) {
	v := T{Id: 123, Name: "Aaron"}

	if !method.Has(v, "Get") {
		t.Fail()
	}

	if method.Has(v, "Method") {
		t.Fail()
	}

	if _, ok := method.Get(v, "Get").(func(T) (int, string)); !ok {
		t.Fail()
	}

	vv := method.Call(v, "Get")
	if _vv, ok := vv[0].(int); !ok || _vv != 123 {
		t.Fail()
	}

	if _vv, ok := vv[1].(string); !ok || _vv != "Aaron" {
		t.Fail()
	}
}
