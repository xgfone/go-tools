package worker

// Job represents the job to be run
type Job struct {
	Payload interface{}
}

// The task interface to handle the job.
type Task interface {
	Handle(Job)
}

// Worker represents the worker that executes the job
type Worker struct {
	workerPool chan chan Job
	jobChannel chan Job
	quit       chan bool
	handler    interface{}
}

// NewWorker creates a new worker.
//
// The worker registers its job channel into workPool to get the job,
// then handle it by handler, either which implements the interface Task,
// or whose type is func(Job).
func NewWorker(workerPool chan chan Job, handler interface{}) *Worker {
	return &Worker{
		workerPool: workerPool,
		jobChannel: make(chan Job),
		quit:       make(chan bool),
		handler:    handler,
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w *Worker) Start() {
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

func (w *Worker) handle(job Job) {
	recovered := true

	defer func() {
		if recovered {
			recover()
		}
	}()

	if h, ok := w.handler.(Task); ok {
		h.Handle(job)
	} else if h, ok := w.handler.(func(Job)); ok {
		h(job)
	} else {
		recovered = false
		panic("The handler type is wrong")
	}
}

// Stop signals the worker to stop listening for work requests.
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
