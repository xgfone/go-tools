package net2

import (
	"testing"
)

func TestGetIP(t *testing.T) {
	if ips, err := GetIP("127.0.0.1"); err != nil || len(ips) != 1 || ips[0] != "127.0.0.1" {
		t.Fail()
	}

	if ips, err := GetIP("lo"); err != nil || ips[0] != "127.0.0.1" {
		t.Fail()
	}
}

func TestGetInterfaceByIP(t *testing.T) {
	if iface, err := GetInterfaceByIP("127.0.0.1"); err != nil || iface != "lo" {
		return t.Fail()
	}
}
