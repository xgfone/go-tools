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

// GetIPs returns the ip list of the network interface name.
//
// If nicName is a valid ip itself, return it directly.
//
// The ip may be a ipv4 or ipv6, which does not include CIDR, but only ip.
func GetIPs(nicName string) (ips []string, err error) {
	if nicName == "" {
		return nil, fmt.Errorf("the NIC name is empty")
	}

	if net.ParseIP(nicName) != nil {
		return []string{nicName}, nil
	}

	var iface *net.Interface
	if iface, err = net.InterfaceByName(nicName); err != nil {
		return
	}

	var addrs []net.Addr
	if addrs, err = iface.Addrs(); err != nil {
		return
	}

	ips = make([]string, len(addrs))
	for i, _len := 0, len(addrs); i < _len; i++ {
		ips[i] = strings.Split(addrs[i].String(), "/")[0]
	}

	return
}

// GetIP is the alias of the GetIPs.
//
// DEPRECATED! Please use GetIPs.
func GetIP(nicName string) (ips []string, err error) { return GetIPs(nicName) }

// GetMac returns the mac address by the ip or interface name.
func GetMac(ipOrNicName string) (mac string, err error) {
	if ip := net.ParseIP(ipOrNicName); ip != nil {
		ifaces, err := net.Interfaces()
		if err != nil {
			return "", err
		}

		ips := ip.String()
		for _, iface := range ifaces {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			for _, addr := range addrs {
				netip := addr.String()
				if index := strings.IndexByte(netip, '/'); index > 0 {
					netip = netip[:index]
				}

				if ips == netip {
					return iface.HardwareAddr.String(), nil
				}
			}
		}
	} else if iface, err := net.InterfaceByName(ipOrNicName); err != nil {
		return "", err
	} else if iface.Name == ipOrNicName {
		return iface.HardwareAddr.String(), nil
	}

	return "", fmt.Errorf("no mac address associated with '%s'", ipOrNicName)
}

// GetNICNameByIP returns the NIC name by the ip.
//
// If ipOrNicName is a valid NIC name, return it.
func GetNICNameByIP(ipOrNicName string) (nicName string, err error) {
	if ip := net.ParseIP(ipOrNicName); ip != nil {
		ifaces, err := net.Interfaces()
		if err != nil {
			return "", err
		}

		ips := ip.String()
		for _, iface := range ifaces {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			for _, addr := range addrs {
				if strings.Split(addr.String(), "/")[0] == ips {
					return iface.Name, nil
				}
			}
		}

		return "", fmt.Errorf("no NIC binding '%s'", ipOrNicName)
	} else if iface, err := net.InterfaceByName(ipOrNicName); err != nil {
		return "", fmt.Errorf("'%s' is the invalid ip or NIC name", ipOrNicName)
	} else {
		return iface.Name, nil
	}
}

// GetInterfaceByIP is the alias of GetNICNameByIP.
//
// DEPRECATED! Please use GetNICNameByIP.
func GetInterfaceByIP(ip string) (nicName string, err error) {
	return GetNICNameByIP(ip)
}

// GetNICNameAndIP returns the NIC name and the ip by the NIC name or ip.
func GetNICNameAndIP(ipOrNicName string) (nicName, ip string, err error) {
	var ips []string
	if net.ParseIP(ipOrNicName) != nil { // For IP
		ip = ipOrNicName
		nicName, err = GetNICNameByIP(ipOrNicName)
	} else if ips, err = GetIPs(ipOrNicName); err == nil { // For NIC name
		if len(ips) == 0 {
			err = fmt.Errorf("no ips bound to the NIC named '%s'", ipOrNicName)
		} else {
			nicName = ipOrNicName
			ip = ips[0]
		}
	}

	return
}

// GetInterfaceAndIP is the alias of GetNICNameAndIP.
//
// DEPRECATED! Please use GetNICNameAndIP.
func GetInterfaceAndIP(ipOrIface string) (iface, ip string, err error) {
	return GetInterfaceAndIP(ipOrIface)
}

// GetAllIPs returns all the ips on the current host.
func GetAllIPs() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	ips = make([]string, len(addrs))
	for i, _len := 0, len(addrs); i < _len; i++ {
		ips[i] = strings.Split(addrs[i].String(), "/")[0]
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
	ips = make([]string, 0, len(addrs))

OUTER:
	for _, addr := range addrs {
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

		ips = append(ips, ip)
	}
	return
}

// IPIsOnHost returns true if the ip is on the host, or returns false.
//
// If the ip is empty or invalid, return false.
func IPIsOnHost(ip string) bool {
	netip := net.ParseIP(strings.TrimSpace(ip))
	if netip == nil {
		return false
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false
	}

	ips := netip.String()
	for _, addr := range addrs {
		if strings.Split(addr.String(), "/")[0] == ips {
			return true
		}
	}

	return false
}

// IsIP returns true if ip is a valid IPv4 or IPv6, or returns false.
func IsIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
