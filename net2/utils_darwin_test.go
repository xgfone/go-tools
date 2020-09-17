// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func TestGetInterfaceAndIP(t *testing.T) {
	iface, ip, err := GetInterfaceAndIP("127.0.0.1")
	if err != nil || iface != "lo0" || ip != "127.0.0.1" {
		t.Error(iface, ip, err)
	}

	iface, ip, err = GetInterfaceAndIP("lo0")
	if err != nil || iface != "lo0" || ip != "127.0.0.1" {
		t.Error(iface, ip, err)
	}
}

func TestGetMac(t *testing.T) {
	mac1, err := GetMac("127.0.0.1")
	if err != nil {
		t.Error(err)
	}

	mac2, err := GetMac("lo0")
	if err != nil {
		t.Error(err)
	}

	if mac1 != mac2 {
		t.Errorf("inconsistent mac address: %s != %s", mac1, mac2)
	}
}
