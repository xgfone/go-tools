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
