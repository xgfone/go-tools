package net2

import "testing"

func TestGetIP(t *testing.T) {
	if ips, err := GetIP("127.0.0.1"); err != nil || len(ips) != 1 || ips[0] != "127.0.0.1" {
		t.Fail()
	}

	if ips, err := GetIP("lo0"); err != nil || ips[0] != "127.0.0.1" {
		t.Fail()
	}
}
