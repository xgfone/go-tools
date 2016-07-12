package server

import (
	"errors"
	"log"
	"net"
	"os"
)

var (
	_logger *log.Logger
)

func init() {
	_logger = log.New(os.Stderr, "[TCP Server] ", log.LstdFlags)
}

func SetLogger(logger *log.Logger) {
	_logger = logger
}

func GetLogger() *log.Logger {
	return _logger
}

type THandle interface {
	Handle(conn *net.TCPConn)
}

// Wrap a panic, only print it, but ignore it.
//
// handle is a function, whose type is `func (*net.TCPConn)`, or a struct, which
// has implemented the interface `THandle`.
func WrapError(conn *net.TCPConn, handle interface{}) {
	yes := true
	defer func() {
		if err := recover(); err != nil {
			_logger.Printf("[Error] Get a error: %v", err)
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

// Start a TCP server and never return. Return a error when returns.
//
// network MUST be "tcp", "tcp4", "tcp6".
// addr is like "host:port", such as "127.0.0.1:8000", and host or port
// may be omitted.
func TCPServerForever(network, addr string, handle interface{}) (e error) {
	ln, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	tcpln, ok := ln.(*net.TCPListener)
	if !ok {
		return errors.New("Must be a TCP Listener")
	}

	_logger.Printf("[Debug] Listen on %v", addr)
	for {
		conn, err := tcpln.AcceptTCP()
		if err != nil {
			e = err
			_logger.Printf("[Error] Failed to AcceptTCP: %v", err)
		} else {
			e = nil
			go WrapError(conn, handle)
		}
	}
	return nil
}
