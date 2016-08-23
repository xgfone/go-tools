package handler_test

import (
	"fmt"

	"github.com/xgfone/go-tools/log/handler"
)

func ExampleTimedRotatingFile() {
	h := handler.NewTimedRotatingFile("test.log")
	n, err := h.Write([]byte("test"))
	if err != nil || n != 4 {
		fmt.Println("Fail")
	} else {
		fmt.Println("Success")
	}

	// Output:
	// Success
}
