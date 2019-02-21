package option

import (
	"testing"
)

func TestOption(t *testing.T) {
	if NONE.IsSome() || NONE.SomeOr(123).(int) != 123 {
		t.Fail()
	}

	if Some(123).IsNone() || Some(123).SomeOr(456).(int) != 123 {
		t.Fail()
	}
}

func TestNamedOption(t *testing.T) {
	if NamedNone("").IsSome() || NamedNone("").SomeOr(123).(int) != 123 {
		t.Fail()
	}

	if NamedSome("", 123).IsNone() || NamedSome("", 123).SomeOr(456).(int) != 123 {
		t.Fail()
	}
}
