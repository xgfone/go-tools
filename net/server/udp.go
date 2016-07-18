package server

import (
	"errors"
	"net"

	"github.com/xgfone/go-tools/pool"
)

type UHandle interface {
	// Handle the request from the client, and return the handled result.
	//
	// buf is the data received from the client.
	// addr is the address of the client.
	Handle(buf []byte, addr *net.UDPAddr) []byte
}

func UDPWithError(conn *net.UDPConn, handle interface{}, buf []byte, addr *net.UDPAddr) {
	yes := true
	defer func() {
		if err := recover(); err != nil {
			_logger.Error("Get a error: %v", err)
			if !yes {
				panic(err)
			}
		}
	}()

	var result []byte
	if handler, ok := handle.(UHandle); ok {
		result = handler.Handle(buf, addr)
	} else if handler, ok := handle.(func([]byte, *net.UDPAddr) []byte); ok {
		result = handler(buf, addr)
	} else {
		yes = false
		panic("Don't support the handler")
	}

	// If returning nil, don't send the response to the client.
	if result == nil {
		return
	}

	if n, err := conn.WriteToUDP(result, addr); err != nil {
		_logger.Error("Failed to send the data to %s: %v", addr, err)
	} else {
		_logger.Debug("Send %v bytes successfully", n)
	}
}

// Start a UDP server and never return. Return an error if returns.
// But there is one exception: if wrap exists and returns true, it returns nil.
// Or continue to execute and never return.
//
// network MUST be "udp", "udp4", "udp6".
// addr is like "host:port", such as "127.0.0.1:8000", and host or port
// may be omitted.
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
func UDPServerForever(network, addr string, size int, handle interface{}, wrap func(*net.UDPConn) bool) error {
	var conn *net.UDPConn
	if _addr, err := net.ResolveUDPAddr(network, addr); err != nil {
		return err
	} else {
		if conn, err = net.ListenUDP(network, _addr); err != nil {
			return err
		}
	}

	defer conn.Close()

	if size < 1 || size > 65536 {
		return errors.New("The size of the buffer is limited between 1 and 65536.")
	}

	if handle == nil && wrap == nil {
		return errors.New("handle and wrap neither be nil.")
	}

	_logger.Info("Listen on %v", addr)

	if wrap != nil {
		if wrap(conn) {
			return nil
		}
	}

	if handle == nil {
		return nil
	}

	_pool := pool.NewBufPool(size)

	for {
		buf := _pool.Get()
		num, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			_logger.Error("Failed to read the UDP data: %v", err)
		} else {
			//go UDPWithError(conn, handle, buf[:num], caddr)
			UDPWithError(conn, handle, buf[:num], caddr)
		}
		_pool.Put(buf)
	}

	// Never execute forever.
	return nil
}
