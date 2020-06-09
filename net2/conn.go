// Copyright 2020 xgfone
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
	"bufio"
	"net"

	"github.com/xgfone/go-tools/v7/io2"
)

// PeekerConn is a union interface with net.Conn and Peeker.
type PeekerConn interface {
	net.Conn
	io2.Peeker
}

// NewPeekerConn is equal to NewPeekerConnSize(conn, 2048).
func NewPeekerConn(conn net.Conn) PeekerConn {
	return NewPeekerConnSize(conn, 2048)
}

// NewPeekerConnSize returns a new PeekerConn with the size of the read buffer.
func NewPeekerConnSize(conn net.Conn, size int) PeekerConn {
	return peekerConn{Conn: conn, buf: bufio.NewReaderSize(conn, size)}
}

var _ net.Conn = peekerConn{}

// Conn wraps net.Conn, and supports the Peek method with a bufio.Reader.
type peekerConn struct {
	net.Conn
	buf *bufio.Reader
}

func (c peekerConn) Read(b []byte) (n int, err error) {
	return c.buf.Read(b)
}

// Peek returns the next n bytes without advancing the reader.
func (c peekerConn) Peek(n int) ([]byte, error) {
	return c.buf.Peek(n)
}
