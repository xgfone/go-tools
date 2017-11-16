package server

import (
	"errors"
	"net"

	"github.com/xgfone/go-tools/log"
	"github.com/xgfone/go-tools/nets"
)

// UHandle is the interface of UDP server handler.
type UHandle interface {
	// Handle the request from the client, and return the handled result.
	//
	// buf is the data received from the client.
	// addr is the address of the client.
	Handle(buf []byte, addr *net.UDPAddr) []byte
}

// UHandleFunc is the type to wrap the function handler to the interface UHandle.
type UHandleFunc (func([]byte, *net.UDPAddr) []byte)

// Handle is the implementation of UHandle.
func (h UHandleFunc) Handle(buf []byte, addr *net.UDPAddr) []byte {
	return h(buf, addr)
}

// UDPWithError wraps a panic, only print it, but ignore it, when to handle a UDP connection.
func UDPWithError(conn *net.UDPConn, handler UHandle, buf []byte, addr *net.UDPAddr) {
	defer func() {
		if err := recover(); err != nil {
			log.ErrorF("Get a error: %v", err)
		}
	}()

	// If returning nil, don't send the response to the client.
	if result := handler.Handle(buf, addr); result != nil {
		if n, err := conn.WriteToUDP(result, addr); err != nil {
			log.ErrorF("Failed to send the data to %s: %v", addr, err)
		} else {
			log.DebugF("Send %v bytes successfully", n)
		}
	}
}

// UDPServerForever starts a UDP server and never return. Return an error if returns.
// But there is one exception: if wrap exists and returns true, it returns nil.
// Or continue to execute and never return.
//
// addr is like "host:port", such as "127.0.0.1:8000", and host or port may be omitted.
// size is the size of the buffer.
//
// Example:
//   type Handler struct {
//   }
//
//   func (h Handler) Handle(buf []byte, addr *net.UDPAddr) []byte {
//   	fmt.Println("Receive %v bytes from %v", len(buf), addr)
//   	return buf
//   }
//
//    err1 := server.UDPServerForever("udp", ":8000", 8192, Handler{}, nil)
//    fmt.Println(err1)
func UDPServerForever(addr string, size int, handle interface{}) error {
	var handler UHandle
	var wrap func(*net.UDPConn) error
	if _handler, ok := handle.(UHandle); ok {
		handler = _handler
	} else if _handler, ok := handle.(func([]byte, *net.UDPAddr) []byte); ok {
		handler = UHandleFunc(_handler)
	} else if _wrap, ok := handle.(func(*net.UDPConn) error); ok {
		wrap = _wrap
	} else {
		panic("Don't support the handler")
	}

	var conn *net.UDPConn
	_addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	if conn, err = net.ListenUDP("udp", _addr); err != nil {
		return err
	}

	defer conn.Close()
	log.DebugF("Listening on %v", addr)

	if wrap != nil {
		return wrap(conn)
	}

	if size < 1 || size > 65536 {
		return errors.New("The size of the buffer is limited between 1 and 65536")
	}
	buf := make([]byte, size)

	for {
		num, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.ErrorF("Failed to read the UDP data: %v", err)
		} else {
			UDPWithError(conn, handler, buf[:num], caddr)
		}
	}

	// Never execute forever.
	// return nil
}

// DialUDP is the same as DialUDPWithAddr, but it joins host and port firstly.
func DialUDP(host, port interface{}) (*net.UDPConn, error) {
	addr := nets.JoinHostPort(host, port)
	return DialUDPWithAddr(addr)
}

// DialUDPWithAddr dials a tcp connection to addr.
func DialUDPWithAddr(addr string) (*net.UDPConn, error) {
	var err error
	_conn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, err
	}

	conn, ok := _conn.(*net.UDPConn)
	if !ok {
		_conn.Close()
		return nil, errors.New("not the udp connection")
	}

	return conn, nil
}
