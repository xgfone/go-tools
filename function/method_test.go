package function

import (
	"testing"
)

type T struct {
	ID   int
	Name string
}

func (t T) Get() (int, string) {
	return t.ID, t.Name
}

func TestMethod(t *testing.T) {
	v := T{ID: 123, Name: "Aaron"}

	if !HasMethod(v, "Get") {
		t.Fail()
	}

	if HasMethod(v, "Method") {
		t.Fail()
	}

	if GetMethod(v, "ID") != nil {
		t.Fail()
	}

	if _, err := CallMethod(v, "ID"); err == nil {
		t.Fail()
	}

	if _, err := CallMethod(v, "Get", 11); err == nil {
		t.Fail()
	}

	if _, ok := GetMethod(v, "Get").(func(T) (int, string)); !ok {
		t.Fail()
	}

	vv, _ := CallMethod(v, "Get")
	if _vv, ok := vv[0].(int); !ok || _vv != 123 {
		t.Fail()
	}

	if _vv, ok := vv[1].(string); !ok || _vv != "Aaron" {
		t.Fail()
	}
}
