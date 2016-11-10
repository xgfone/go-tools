package server_test

import (
	"fmt"
	"net"

	"github.com/xgfone/go-tools/net/server"
)

func ExampleUDPServerForever() {
	// Here use a function as the handler. You also use a struct which implements
	// the interface THandle.
	err1 := server.UDPServerForever("udp", ":9000", 9120, func(buf []byte, addr *net.UDPAddr) []byte {
		fmt.Printf("Receive the data from %v: %v\n", addr, buf)
		return buf
	})
	fmt.Println(err1)
}
