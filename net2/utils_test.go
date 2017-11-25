package net2

import "testing"

func TestJoinHostPort(t *testing.T) {
	result := "127.0.0.1:8000"

	if JoinHostPort("127.0.0.1", "8000") != result {
		t.Fail()
	}

	if JoinHostPort("127.0.0.1", 8000) != result {
		t.Fail()
	}

	if JoinHostPort([]byte("127.0.0.1"), 8000) != result {
		t.Fail()
	}

	if JoinHostPort([]byte("127.0.0.1"), "8000") != result {
		t.Fail()
	}
}

func TestGetAllIPs(t *testing.T) {
	if ips, err := GetAllIPs(); err != nil || len(ips) == 0 {
		t.Fail()
	}
}

func TestIPIsOnHost(t *testing.T) {
	if !IPIsOnHost("127.0.0.1") {
		t.Fail()
	}
}
