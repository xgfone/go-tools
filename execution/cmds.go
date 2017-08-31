package execution

// Cmd is an once executioin.
var Cmd = NewOnceExecution()

// Output is short for Cmd.Output.
func Output(args []string) (output string, err error) {
	return Cmd.Output(args)
}

// Execute is short for Cmd.Execute.
func Execute(args []string) (err error) {
	return Cmd.Execute(args)
}

// ErrOutput is short for Cmd.ErrOutput.
func ErrOutput(args []string) (output string, err error) {
	return Cmd.ErrOutput(args)
}

// NewOnceExecution returns a new execution, which will be executed once.
//
// Notice: don't modify the field of the returned execution, Count and Interval.
func NewOnceExecution() *Execution {
	return new(Execution)
}

// NewRetryExecution returns a new retry execution, which will be executed
// repeatedly until success or some times.
//
// retry is the times to be executed. If 0, it will be executed once, no retry
// when failure. interval is the interval time between two executions, the unit
// of which is millisecond.
//
// Notice: don't modify the field of the returned execution, Count and Interval.
func NewRetryExecution(retry int, interval int) *Execution {
	if retry < 0 || interval < 0 {
		panic("the argument can not be a negative number")
	}

	return &Execution{
		Count:    retry,
		Interval: interval,
	}
}

// NewRedoExecution returns a new redo execution, which will be executed
// repeatedly until failure or some times.
//
// redo is the times to be executed. If 0, it will be executed once, no redo
// when success. interval is the interval time between two executions, the unit
// of which is millisecond.
//
// Notice: don't modify the field of the returned execution, Count and Interval.
func NewRedoExecution(redo int, interval int) *Execution {
	if redo < 0 || interval < 0 {
		panic("the argument can not be a negative number")
	}

	return &Execution{
		Count:    -redo,
		Interval: interval,
	}
}
