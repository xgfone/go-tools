package worker

// Dispatcher dispatches the task to the workers.
type Dispatcher struct {
	// The number of the worker.
	WorkerNum int

	// The job handler.
	Handler Task

	// A pool of workers channels that are registered with the dispatcher
	workerPool chan chan interface{}

	workers []*worker
	quit    chan bool
	running bool
	async   bool
}

// NewDispatcher creates a new Dispatcher.
func NewDispatcher(maxWorkers int, handler Task) *Dispatcher {
	if maxWorkers < 1 {
		panic("The worker number must be greater than 0")
	}

	return &Dispatcher{
		WorkerNum:  maxWorkers,
		Handler:    handler,
		workerPool: make(chan chan interface{}, maxWorkers),
		quit:       make(chan bool),
	}
}

// SetAsync sets whether to dispatch the task asynchronously.
//
// When true, it's asynchronous, and don't block the Dispatcher when the job
// queue is full. The default is synchronous.
func (d *Dispatcher) SetAsync(b bool) *Dispatcher {
	d.async = b
	return d
}

// Dispatch starts the dispatcher and let the workers to handle the job.
//
// Notice: this method is not thread-safe.
func (d *Dispatcher) Dispatch(jobQueue chan interface{}) {
	if d.running {
		panic("This dispatcher has been running")
	}

	d.workers = make([]*worker, d.WorkerNum)

	// starting n number of workers
	for i := 0; i < d.WorkerNum; i++ {
		worker := newWorker(d.workerPool, d.Handler)
		d.workers[i] = worker
		worker.Start()
	}

	go d.dispatch(jobQueue)
	d.running = true
}

// Stop the dispatcher and all the workers.
//
// Notice: this method is not thread-safe.
func (d *Dispatcher) Stop() {
	d.quit <- true
	for _, worker := range d.workers {
		worker.Stop()
	}
}

func (d *Dispatcher) dispatch(jobQueue chan interface{}) {
	run := func(job interface{}) {
		// try to obtain a worker job channel that is available.
		// this will block until a worker is idle.
		jobChannel := <-d.workerPool

		// dispatch the job to the worker job channel.
		jobChannel <- job
	}

	for {
		select {
		case job := <-jobQueue:
			// a job request has been received
			if d.async {
				go run(job)
			} else {
				run(job)
			}
		case <-d.quit:
			d.running = false
			return
		}
	}
}
