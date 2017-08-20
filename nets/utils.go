package nets

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

// GetIP returns the ip of the network interface name.
//
// If the argument iname is a valid ip itself, return it directly.
//
// The ip may be a ipv4 or ipv6, which does not include CIDR, but only ip.
func GetIP(iname string) (ips []string, err error) {
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

	if len(ips) == 0 {
		err = fmt.Errorf("not found the ip of %s", iname)
	}
	return
}
