package nets_test

import (
	"testing"

	"github.com/xgfone/go-tools/nets"
)

func TestJoinHostPort(t *testing.T) {
	result := "127.0.0.1:8000"

	if nets.JoinHostPort("127.0.0.1", "8000") != result {
		t.Fail()
	}

	if nets.JoinHostPort("127.0.0.1", 8000) != result {
		t.Fail()
	}

	if nets.JoinHostPort([]byte("127.0.0.1"), 8000) != result {
		t.Fail()
	}

	if nets.JoinHostPort([]byte("127.0.0.1"), "8000") != result {
		t.Fail()
	}
}
