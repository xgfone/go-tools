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

import (
	"fmt"
	"net"
	"strings"
)

// JoinHostPort is same as net.JoinHostPort, but it receives the arguments of
// any type, not only string.
//
// Recommend: Only use string or []byte as the type of host, and string or
// integer as that of port.
func JoinHostPort(host, port interface{}) string {
	if _host, ok := host.([]byte); ok {
		host = string(_host)
	}
	return net.JoinHostPort(fmt.Sprintf("%v", host), fmt.Sprintf("%v", port))
}

func getIPByName(iname string, empty bool) (ips []string, err error) {
	if len(iname) == 0 {
		return nil, fmt.Errorf("the parameter is empty")
	}

	if ip := net.ParseIP(iname); ip != nil {
		return []string{iname}, nil
	}

	var _interface *net.Interface
	if _interface, err = net.InterfaceByName(iname); err != nil {
		return
	}

	var addrs []net.Addr
	if addrs, err = _interface.Addrs(); err != nil {
		return
	}
	for _, addr := range addrs {
		ips = append(ips, strings.Split(addr.String(), "/")[0])
	}

	if empty && len(ips) == 0 {
		err = fmt.Errorf("not found the ip of %s", iname)
	}
	return
}

// GetIP returns the ip of the network interface name.
//
// If the argument iname is a valid ip itself, return it directly.
//
// The ip may be a ipv4 or ipv6, which does not include CIDR, but only ip.
func GetIP(iname string) (ips []string, err error) {
	return getIPByName(iname, true)
}

// GetInterfaceByIP returns the interface name bound the ip.
func GetInterfaceByIP(ip string) (iface string, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	ip = strings.ToLower(ip)
	for _, iface := range ifaces {
		_ips, err := getIPByName(iface.Name, false)
		if err != nil {
			return "", err
		}
		for _, _ip := range _ips {
			if _ip == ip {
				return iface.Name, nil
			}
		}
	}

	return "", fmt.Errorf("not found the interface bound '%s'", ip)
}

// GetInterfaceAndIP returns the interface name and the ip
// by the interface name or ip.
func GetInterfaceAndIP(ipOrIface string) (iface, ip string, err error) {
	var ips []string
	if IsIP(ipOrIface) {
		ip = ipOrIface
		iface, err = GetInterfaceByIP(ipOrIface)
	} else if ips, err = GetIP(ipOrIface); err == nil {
		if len(ips) == 0 {
			err = fmt.Errorf("no ip on the interface '%s'", ipOrIface)
		} else {
			iface = ipOrIface
			ip = ips[0]
		}
	}
	return
}

// GetAllIPs returns all the ips on the current host.
func GetAllIPs() (ips []string, err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	var _ips []string
	for _, iface := range ifaces {
		if _ips, err = getIPByName(iface.Name, false); err != nil {
			return
		}
		for _, ip := range _ips {
			ips = append(ips, ip)
		}
	}

	return
}

// GetAllNetIPs returns the ips with the subnet mask, such as
//
//   127.0.0.1/8
//   10.218.0.26/32
//   ::1/128
//   fe80::1/64
//   fe80::ac48:81ff:fe8d:e8b7/64
//
func GetAllNetIPs() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	var ip string
	ips = make([]string, len(addrs))

OUTER:
	for i, addr := range addrs {
		if ip = addr.String(); strings.IndexByte(ip, '/') < 0 {
			if strings.IndexByte(ip, ':') < 0 {
				ip += "/32"
			} else {
				ip += "/128"
			}
		}

		for _, _ip := range ips {
			if _ip == ip {
				continue OUTER
			}
		}

		ips[i] = ip
	}
	return
}

// IPIsOnHost returns true if the ip is on the host, or returns false.
func IPIsOnHost(ip string) bool {
	ip = strings.ToLower(ip)
	ips, _ := GetAllIPs()
	for _, _ip := range ips {
		if _ip == ip {
			return true
		}
	}
	return false
}

// IsIP returns true if ip is a valid IPv4 or IPv6, or returns false.
func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
