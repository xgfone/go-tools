package tb_test

import (
	"fmt"

	"github.com/xgfone/go-tools/tb"
)

func ExampleTB() {
	// Get a token from the bucket per second.
	_tb := tb.NewTB(1)
	_tb.Start()
	go func(_tb *tb.TB) {
		for {
			_tb.Get()
			fmt.Println("Get a token") // You can see it prints once per second.
		}
	}(_tb)

	for {
	}
}
