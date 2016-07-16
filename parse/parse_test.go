package parse_test

import (
	"testing"

	"github.com/xgfone/go-tools/parse"
)

func TestParse(t *testing.T) {
	// Test bool
	if parse.ToBool("t") != true {
		t.Fail()
	}

	if parse.ToB("aaa") != false {
		t.Fail()
	}

	// Test Float
	if parse.ToF32("1.2") != 1.2 {
		t.Fail()
	}

	if parse.ToF64("11aa") != 0.0 {
		t.Fail()
	}

	// Test Int
	if parse.ToInt("123", 10) != 123 {
		t.Fail()
	}

	if parse.ToI("-FF", 16) != -0xFF {
		t.Fail()
	}

	// Test Uint
	if parse.ToUint("FF", 16) != 0xFF {
		t.Fail()
	}

	if parse.ToU("123", 10) != 123 {
		t.Fail()
	}
}
