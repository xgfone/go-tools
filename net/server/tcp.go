package server

import (
	"net"

	"github.com/xgfone/go-tools/pool"
)

type THandle interface {
	Handle(conn *net.TCPConn)
}

// Wrap a panic, only print it, but ignore it.
//
// handle is a function, whose type is `func (*net.TCPConn)`, or a struct, which
// has implemented the interface `THandle`.
// wrap is the wrapper of *net.TCPConn, which sets the socket connection, and
// is used by TCPServerForever. In general, it is nil.
func TCPWrapError(conn *net.TCPConn, handle interface{}) {
	yes := true
	defer func() {
		if err := recover(); err != nil {
			_logger.Error("Get a error: %v", err)
			if !yes {
				panic(err)
			}
		}
	}()
	defer conn.Close()

	if handler, ok := handle.(THandle); ok {
		handler.Handle(conn)
	} else if handler, ok := handle.(func(*net.TCPConn)); ok {
		handler(conn)
	} else {
		yes = false
		panic("Don't support the handler")
	}
}

// Start a TCP server and never return. Return an error if returns.
//
// network MUST be "tcp", "tcp4", "tcp6".
// addr is like "host:port", such as "127.0.0.1:8000", and host or port may be omitted.
// size is the number of the pool. If it's 0, it's infinite.
// handle is the handler to handle the connection came from the client.
// handle is either a function whose type is func(*net.TCPConn), or a struct
// which implements the interface, THandle.
func TCPServerForever(network, addr string, size int, handle interface{}) error {
	var ln *net.TCPListener
	if _addr, err := net.ResolveTCPAddr(network, addr); err != nil {
		return err
	} else {
		if ln, err = net.ListenTCP(network, _addr); err != nil {
			return err
		}
	}

	defer ln.Close()

	var gopool *pool.GoPool
	if size > 0 {
		gopool = pool.NewGoPool()
		gopool.SetMaxLimit(uint(size))
	}

	_logger.Info("Listen on %v", addr)

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			_logger.Error("Failed to AcceptTCP: %v", err)
		} else {
			_logger.Debug("Get a connection from %v", conn.RemoteAddr())
			if gopool == nil {
				go TCPWrapError(conn, handle)
			} else {
				if err := gopool.Go(TCPWrapError, conn, handle); err != nil {
					_logger.Error("Failed to run goroutine: %v", err)
				}
			}
		}
	}

	// Never execute forever.
	return nil
}
