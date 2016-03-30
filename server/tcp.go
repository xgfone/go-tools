// ```go
// // TCP Server
// package main
//
// import (
// 	"fmt"
// 	"github.com/xgfone/go-tools/server"
// 	"io"
// 	"net"
// )
//
// func handle(conn *net.TCPConn) {
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				fmt.Println("Conn broke off")
// 				return
// 			}
// 			fmt.Println(err)
// 		} else {
// 			fmt.Printf("Receive %v bytes: %v\n", n, string(buf[:n]))
// 		}
// 	}
// }
//
// func main() {
// 	err := server.TCPServerForever("tcp", ":8000", handle)
// 	fmt.Println(err)
// }
// ```
//
// ```go
// // TCP Client
// package main
//
// import (
// 	"fmt"
// 	"net"
// 	"os"
// 	"time"
// )
//
// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:8000")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	tcp, ok := conn.(*net.TCPConn)
// 	for ok {
// 		n, err := tcp.Write([]byte("hello"))
// 		if err != nil {
// 			fmt.Printf("Error: %v\n", err)
// 			return
// 		}
// 		fmt.Printf("Write %v bytes\n", n)
// 		time.Sleep(time.Second)
// 	}
// }
// ```

package server

import (
	"errors"
	"fmt"
	"net"
)

type THandle func(*net.TCPConn)

func WrapError(handle THandle, conn *net.TCPConn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Get a error: %v\n", err)
		}
	}()
	handle(conn)
	conn.Close()
}

func TCPServerForever(network, addr string, handle THandle) (e error) {
	ln, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	tcpln, ok := ln.(*net.TCPListener)
	if !ok {
		return errors.New("Must be a TCP Listener")
	}

	fmt.Printf("Listen on %v\n", addr)
	for {
		conn, err := tcpln.AcceptTCP()
		if err != nil {
			e = err
			fmt.Printf("Failed to AcceptTCP: %v\n", err)
		} else {
			e = nil
			go WrapError(handle, conn)
		}
	}
	return nil
}
