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

func TestIsIP(t *testing.T) {
	if !IsIP("1.2.3.4") || !IsIP("fe80::acf4:ffff:feb7:bb24") {
		t.Fail()
	}

	if IsIP("1.2.3.4.5") || IsIP("fe80::acf4::ffff:feb7:bb24") {
		t.Fail()
	}
}
