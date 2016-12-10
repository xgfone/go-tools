package execution_test

import "github.com/xgfone/go-tools/execution"

func ExampleExecution_Execute() {
	e := execution.Execution{Retry: 1, Interval: 1000}
	err := e.Execute([]string{"ls", "."})
}

func ExampleExecution_Output() {
	e := execution.Execution{Retry: 1, Interval: 1000}
	out, err := e.Output([]string{"ls", "."})
}

func ExampleExecution_ErrOutput() {
	e := execution.Execution{Retry: 1, Interval: 1000}
	out, err := e.ErrOutput([]string{"ls", "."})
}
