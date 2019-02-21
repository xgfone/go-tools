package types

import "testing"

func TestConverter(t *testing.T) {
	c := NewConverter(Int)
	if err := c.Scan("123"); err != nil {
		t.Error(err)
	} else if v := c.Value().(int); v != 123 {
		t.Error(v)
	}
}
