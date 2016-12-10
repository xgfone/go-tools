package server_test

import (
	"fmt"
	"net"

	"github.com/xgfone/go-tools/nets/server"
)

func ExampleUDPServerForever() {
	// Here use a function as the handler. You also use a struct which implements
	// the interface UHandle.
	err1 := server.UDPServerForever("udp", ":9000", 9120, func(buf []byte, addr *net.UDPAddr) []byte {
		fmt.Printf("Receive the data from %v: %v\n", addr, buf)
		return buf
	})
	fmt.Println(err1)
}

func ExampleDialUDP() {
	if conn, err := server.DialUDP("223.5.5.5", 53); err != nil {
		fmt.Printf("ERR\n")
	} else {
		fmt.Printf("OK\n")
		conn.Close()
	}

	// Output:
	// OK
}

func ExampleDialUDPWithAddr() {
	if conn, err := server.DialUDPWithAddr("223.5.5.5:53"); err != nil {
		fmt.Printf("ERR\n")
	} else {
		fmt.Printf("OK\n")
		conn.Close()
	}

	// Output:
	// OK
}
