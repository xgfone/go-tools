package execution

type RetryExecution struct {
	*Execution
}

func NewRetryExecution(retry int, interval int) *RetryExecution {
	r := new(RetryExecution)
	r.SetRetryNum(retry)
	r.Interval = interval
	return r
}

func (r *RetryExecution) SetRetryNum(n int) {
	if n < 0 {
		panic("the argument must be a positive number")
	}
	r.Count = n
}
