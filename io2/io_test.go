package io2

import (
	"bytes"
	"fmt"
	"io"
)

func ExampleReadN() {
	s := "1234567890"
	rbuf := bytes.NewBufferString(s)
	if v, err := ReadN(rbuf, 9); err != nil || string(v) != s[:9] {
		fmt.Println("Error")
	} else {
		fmt.Println("OK")
	}

	rbuf = bytes.NewBufferString(s)
	if v, err := ReadN(rbuf, 0); err != nil || string(v) != s {
		fmt.Println("Error")
	} else {
		fmt.Println("OK")
	}

	rbuf = bytes.NewBufferString(s)
	if v, err := ReadN(rbuf, 11); err != io.EOF || string(v) != s {
		fmt.Println("Error")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
	// OK
	// OK
}
