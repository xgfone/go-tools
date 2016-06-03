package exception_test

import (
	"testing"

	"github.com/xgfone/go-tools/exception"
)

func TestException(t *testing.T) {
	s1 := "aa.bb.cc"
	s2 := "aa.bb"
	s3 := "dd.ee"

	e1 := exception.New(s1, "Exception1")
	e2 := exception.New(s2, "Exception2")
	e3 := exception.New(s3, "Exception3")

	// fmt.Println(e1, e2, e3, e1.IsChild(e2))

	if !e1.IsChild(e2) {
		t.Fail()
	}

	if e1.IsChild(e3) {
		t.Fail()
	}

	if e2.IsChild(e1) {
		t.Fail()
	}

	if e2.IsChild(e3) {
		t.Fail()
	}

	if e3.IsChild(e1) {
		t.Fail()
	}

	if e3.IsChild(e2) {
		t.Fail()
	}

	if !e1.IsSame(e1) {
		t.Fail()
	}

	if e1.IsSame(e2) {
		t.Fail()
	}
}
