// The simple TCP/UDP server.
package server

import "net"

type THandle interface {
	Handle(conn *net.TCPConn)
}

// Wrap the function handler to the interface THandle.
type THandleFunc (func(*net.TCPConn))

func (h THandleFunc) Handle(conn *net.TCPConn) {
	h(conn)
}

// Wrap a panic, only print it, but ignore it.
func TCPWrapError(conn *net.TCPConn, handler THandle) {
	defer func() {
		if err := recover(); err != nil {
			_logger.Error("Get a error: %v", err)
		}

		if conn != nil {
			conn.Close()
		}
	}()

	handler.Handle(conn)
}

// Start a TCP server and never return. Return an error if returns.
//
// network MUST be "tcp", "tcp4", "tcp6".
// addr is like "host:port", such as "127.0.0.1:8000", and host or port may be omitted.
// size is the number of the pool. If it's 0, it's infinite.
// handle is the handler to handle the connection came from the client.
// handle is either a function whose type is func(*net.TCPConn), or a struct
// which implements the interface, THandle. Of course, you may wrap it by THandleFunc.
func TCPServerForever(network, addr string, handle interface{}) error {
	var handler THandle
	if _handler, ok := handle.(THandle); !ok {
		handler = _handler
	} else if _handler, ok := handle.(func(*net.TCPConn)); ok {
		handler = THandleFunc(_handler)
	} else {
		panic("Don't support the handler")
	}

	var ln *net.TCPListener
	if _addr, err := net.ResolveTCPAddr(network, addr); err != nil {
		return err
	} else {
		if ln, err = net.ListenTCP(network, _addr); err != nil {
			return err
		}
	}

	defer ln.Close()

	_logger.Info("Listening on %v", addr)

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			_logger.Error("Failed to AcceptTCP: %v", err)
		} else {
			_logger.Debug("Get a connection from %v", conn.RemoteAddr())
			go TCPWrapError(conn, handler)
		}
	}

	// Never execute forever.
	return nil
}
