package execution

type RedoExecution struct {
	*Execution
}

func NewRedoExecution(redo int, interval int) *RedoExecution {
	r := new(RedoExecution)
	r.SetRedoNum(redo)
	r.Interval = interval
	return r
}

func (r *RedoExecution) SetRedoNum(n int) {
	if n < 0 {
		panic("the argument must be a positive number")
	}
	r.Count = -n
}
