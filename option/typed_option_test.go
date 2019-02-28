package option

import (
	"testing"
)

func TestBoolOption(t *testing.T) {
	b := NewBoolOption(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan("true"); err != nil {
		t.Error(err)
	} else if !b.IsBool() || !b.Bool() {
		t.Error(b.Value())
	}
}

func TestInt8Option(t *testing.T) {
	b := NewInt8Option(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan("123"); err != nil {
		t.Error(err)
	} else if !b.IsInt8() || b.Int8() != 123 {
		t.Error(b.Value())
	}
}

func TestFloat64Option(t *testing.T) {
	b := NewFloat64Option(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan("1.2"); err != nil {
		t.Error(err)
	} else if !b.IsFloat64() || b.Float64() != 1.2 {
		t.Error(b.Value())
	}
}

func TestStringOption(t *testing.T) {
	b := NewStringOption(None())

	if b.IsSome() {
		t.Fail()
	}

	if err := b.Scan(123); err != nil {
		t.Error(err)
	} else if !b.IsString() || b.Str() != "123" {
		t.Error(b.Value())
	}
}

func TestInterface(t *testing.T) {
	opts := []Option{
		NewBoolOption(None()),
		NewInt64Option(None()),
		NewFloat64Option(None()),
		NewStringOption(None()),
	}
	values := []interface{}{"true", "123", "1.2", 456}

	for i, opt := range opts {
		if err := opt.Scan(values[i]); err != nil {
			t.Error(err)
		} else if opt.Value() == nil {
			t.Fail()
		}
	}
}

func TestNamedTypedOption(t *testing.T) {
	opt := NewNamedOption("", NewBoolOption(nil))
	if err := opt.Scan("on"); err != nil {
		t.Error(err)
	} else if !opt.IsBool() || !opt.Bool() {
		t.Error(opt)
	}
}
