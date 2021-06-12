// Copyright 2021 xgfone
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
	"net"
	"strings"
)

// IPIsOnHost returns true if the ip is on the host, or returns false.
//
// If the ip is empty or invalid, return false.
func IPIsOnHost(ip string) bool {
	netip := net.ParseIP(strings.TrimSpace(ip))
	if netip == nil {
		return false
	}

	ip = netip.String()
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if strings.Split(addr.String(), "/")[0] == ip {
			return true
		}
	}

	return false
}
