package server

import (
	"fmt"
	"time"

	"github.com/xgfone/go-tools/lifecycle"
)

func ExampleForeverAndShutdown() {
	lifecycle.Register(func() {
		fmt.Println("App clean")
	})

	go func() {
		for !IsShutdowned() {
			fmt.Println("Running")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		Shutdown()
	}()

	RunForever()
	fmt.Println("Program exits")

	// Output:
	// Running
	// Running
	// Running
	// App clean
	// Program exits
}
