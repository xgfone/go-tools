package net2

import (
	"net"
)

// ListenUDP listens UDP on addr, then returns a UDP connection.
func ListenUDP(addr string) (*net.UDPConn, error) {
	_addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	return net.ListenUDP("udp", _addr)
}

// DialUDP dials a UDP connection to host:port.
func DialUDP(host, port interface{}) (*net.UDPConn, error) {
	return DialUDPByAddr(JoinHostPort(host, port))
}

// DialUDPByAddr dials a UDP connection to addr.
func DialUDPByAddr(addr string) (*net.UDPConn, error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}
	return conn.(*net.UDPConn), nil
}
