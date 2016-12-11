package execution_test

import (
	"fmt"
	"sync"

	"github.com/xgfone/go-tools/execution"
)

var lock = new(sync.Mutex)

func ExampleExecution_Execute() {
	e := execution.Execution{Count: 1, Interval: 1000, IsLock: true}
	e.SetMutex(lock)
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
	e := execution.Execution{Count: -1, Interval: 1000, IsLock: true}
	e.SetMutex(lock)
	_, err := e.Output([]string{"ls", "."})
	if err == nil { // Notice: we run it until failed
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// ERROR
}

func ExampleExecution_ErrOutput() {
	e := execution.Execution{Count: 1, Interval: 1000, IsLock: true}
	e.SetMutex(lock)
	_, err := e.ErrOutput([]string{"ls", "."})
	if err != nil {
		fmt.Println("ERROR")
	} else {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
