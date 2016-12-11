package execution_test

import (
	"fmt"

	"github.com/xgfone/go-tools/execution"
)

func ExampleExecution_Execute() {
	e := execution.Execution{Count: 1, Interval: 1000}
	err := e.Execute([]string{"ls", "."})
	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleExecution_Output() {
	e := execution.Execution{Count: 1, Interval: 1000}
	_, err := e.Output([]string{"ls", "."})
	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleExecution_ErrOutput() {
	e := execution.Execution{Count: 1, Interval: 1000}
	_, err := e.ErrOutput([]string{"ls", "."})
	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
