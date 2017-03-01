// Package worker is a worker pool with the dispatcher based on channel.
//
// Refer to http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang.
package worker

// Task is an interface to handle the job.
type Task interface {
	// The argument is the job object.
	Handle(interface{})
}

// FuncTask converts a function to Task.
type FuncTask func(interface{})

// Handle implements the interface Task.
func (f FuncTask) Handle(job interface{}) {
	f(job)
}

// worker represents the worker that executes the job
type worker struct {
	workerPool chan chan interface{}
	jobChannel chan interface{}
	quit       chan bool
	handler    Task
}

// newWorker creates a new worker.
//
// The worker registers its job channel into workPool to get the job,
// then handle it by handler.
func newWorker(workerPool chan chan interface{}, handler Task) *worker {
	return &worker{
		workerPool: workerPool,
		jobChannel: make(chan interface{}),
		quit:       make(chan bool),
		handler:    handler,
	}
}

// Start starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w *worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.workerPool <- w.jobChannel

			select {
			case job := <-w.jobChannel:
				// we have received a work request.
				w.handle(job)
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

func (w *worker) handle(job interface{}) {
	defer recover()
	w.handler.Handle(job)
}

// Stop signals the worker to stop listening for work requests.
func (w *worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
