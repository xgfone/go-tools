package execution

import (
	"context"
	"fmt"
)

func ExampleExecute() {
	if err := Execute(context.TODO(), "ls", "."); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleExecutes() {
	if err := Executes(context.TODO(), []string{"ls"}); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleOutput() {
	if _, err := Output(context.TODO(), "ls", "."); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleOutputs() {
	if _, err := Outputs(context.TODO(), []string{"ls"}); err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
