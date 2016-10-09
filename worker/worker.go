// A worker pool with the dispatcher based on channel.
package worker

// The task interface to handle the job.
type Task interface {
	// The argument is the job object.
	Handle(interface{})
}

// Worker represents the worker that executes the job
type worker struct {
	workerPool chan chan interface{}
	jobChannel chan interface{}
	quit       chan bool
	handler    interface{}
}

// NewWorker creates a new worker.
//
// The worker registers its job channel into workPool to get the job,
// then handle it by handler, either which implements the interface Task,
// or whose type is func(Job).
func newWorker(workerPool chan chan interface{}, handler interface{}) *worker {
	return &worker{
		workerPool: workerPool,
		jobChannel: make(chan interface{}),
		quit:       make(chan bool),
		handler:    handler,
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
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
	recovered := true

	defer func() {
		if recovered {
			recover()
		}
	}()

	if h, ok := w.handler.(Task); ok {
		h.Handle(job)
	} else if h, ok := w.handler.(func(interface{})); ok {
		h(job)
	} else {
		recovered = false
		panic("The handler type is wrong")
	}
}

// Stop signals the worker to stop listening for work requests.
func (w *worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
