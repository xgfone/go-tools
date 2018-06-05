package net2

import (
	"fmt"
)

func ExampleGetMacByInterface() {
	_, err := GetMacByInterface("lo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleGetMacByIP() {
	_, err := GetMacByIP("127.0.0.1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
