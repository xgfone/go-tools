package server_test

import (
	"fmt"
	"io"
	"net"

	"github.com/xgfone/go-tools/net/server"
)

func ExampleTCPServerForever() {
	// Here use a function as the handler. You also use a struct which implements
	// the interface THandle.
	err1 := server.TCPServerForever("tcp", ":8000", func(conn *net.TCPConn) {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Conn broke off")
					return
				}
				fmt.Println(err)
			} else {
				fmt.Printf("Receive %v bytes: %v\n", n, string(buf[:n]))
			}
		}
	})
	fmt.Println(err1)
}
