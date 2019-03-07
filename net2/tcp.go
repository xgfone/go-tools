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
)

// TCPServerForever starts a TCP server. If starting successfully, never return.
func TCPServerForever(addr string, handler func(*net.TCPConn)) error {
	_addr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}

	ln, err := net.ListenTCP("tcp", _addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Printf("AcceptTCP get an error: %v\n", err)
		} else {
			go handler(conn)
		}
	}

	// Never execute forever.
	// return nil
}

// DialTCP dials a TCP connection to host:port.
func DialTCP(host, port interface{}) (*net.TCPConn, error) {
	return DialTCPByAddr(JoinHostPort(host, port))
}

// DialTCPByAddr dials a TCP connection to addr.
func DialTCPByAddr(addr string) (*net.TCPConn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn.(*net.TCPConn), nil
}
