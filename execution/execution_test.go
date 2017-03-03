package execution_test

import (
	"fmt"
	"sync"

	"github.com/xgfone/go-tools/execution"
)

var lock = new(sync.Mutex)

func ExampleExecution_Execute() {
	Retry := execution.Execution{Count: 1, Interval: 1000}
	Retry.SetMutex(lock)
	err := Retry.Execute([]string{"ls", "."})
	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleExecution_Output() {
	Redo := execution.Execution{Count: -1, Interval: 1000}
	Redo.SetMutex(lock)
	_, err := Redo.Output([]string{"ls", "."})
	if err == nil { // Notice: we run it until failed
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// ERROR
}

func ExampleExecution_ErrOutput() {
	Retry := execution.Execution{Count: 1, Interval: 1000}
	Retry.SetMutex(lock)
	_, err := Retry.ErrOutput([]string{"ls", "."})
	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
