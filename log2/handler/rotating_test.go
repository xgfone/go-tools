package handler

import (
	"fmt"
)

func ExampleSizedRotatingFile() {
	h := NewSizedRotatingFile("test_rotatingfile.log", 1024, 3)
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
