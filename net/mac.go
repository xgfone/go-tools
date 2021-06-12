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

package net

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// NormalizeMac normalizes the mac. But return "" if the mac is a invalid mac.
func NormalizeMac(mac string) string {
	macs := strings.Split(mac, ":")
	if len(macs) != 6 {
		return ""
	}

	for i, m := range macs {
		v, err := strconv.ParseUint(m, 16, 8)
		if err != nil {
			return ""
		}
		macs[i] = fmt.Sprintf("%02x", v)
	}

	return strings.Join(macs, ":")
}

// GetMacByInterface returns the MAC of the interface iface.
func GetMacByInterface(iface string) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		if i.Name == iface {
			return i.HardwareAddr.String(), nil
		}
	}

	return "", fmt.Errorf("no interface '%s'", iface)
}

// GetMacByIP returns the MAC of the interface to which the ip is bound.
func GetMacByIP(ip string) (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	ip = strings.ToLower(ip)
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			_ip := strings.Split(addr.String(), "/")[0]
			if _ip == ip {
				return iface.HardwareAddr.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no mac about '%s'", ip)
}
