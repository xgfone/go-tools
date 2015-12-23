package server

import (
	"fmt"
	"net"
)

type App interface {
	Handle(*net.TCPConn)
}

func WrapError(app App, conn *net.TCPConn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Get a error: %v\n", err)
		}
	}()
	app.Handle(conn)
}

func TCPServerForever(net, addr string, app App) (e error) {
	ln, err := net.Listen(net, addr)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			e = err
			fmt.Printf("Failed to AcceptTCP: %v\n", err)
		} else {
			e = nil
			go WrapError(app, conn)
		}
	}
	return nil
}
