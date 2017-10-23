package handler

import (
	"fmt"
)

func ExampleTimedRotatingFile() {
	h := NewTimedRotatingFile("test.log")
	defer h.Close()
	n, err := h.Write([]byte("test"))
	if err != nil || n != 4 {
		fmt.Println("Fail")
	} else {
		fmt.Println("Success")
	}

	// Output:
	// Success
}

func ExampleRotatingFile() {
	h := NewRotatingFile("test_rotatingfile.log", 1024, 3)
	defer h.Close()
	for i := 1; i < 1000; i++ {
		data := fmt.Sprintf("test the RotatingFile %d\n", i)
		if _, err := h.WriteString(data); err != nil {
			fmt.Printf("Fail: %s\n", err)
			return
		}
	}
	fmt.Println("Success")

	// Output:
	// Success
}
