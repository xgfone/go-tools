// Package worker is a worker pool with the dispatcher based on channel.
package worker

import (
	"context"
)

// Task is an interface to handle the job.
type Task interface {
	// The argument is the job object.
	Handle(job interface{})
}

// TaskFunc converts a function to Task.
type TaskFunc func(interface{})

// Handle implements the interface Task.
func (f TaskFunc) Handle(job interface{}) {
	f(job)
}

// PutJob puts the job into the job queue.
//
// If the queue is full, it will be blocked.
func PutJob(jobQueue chan<- interface{}, job interface{}) {
	jobQueue <- job
}

// TryPutJob puts the job into the job queue and returns true.
//
// If the queue is full, it will pass it and return false.
func TryPutJob(jobQueue chan<- interface{}, job interface{}) bool {
	select {
	case jobQueue <- job:
		return true
	default:
		return false
	}
}

// Dispatch starts a worker pool and dispatches the task to it.
//
// Notice: you maybe cancel the worker pool by the context.
func Dispatch(cxt context.Context, workerNum int, jobQueue <-chan interface{},
	handler Task) {

	// Call the handler to handle the job.
	handleJob := func(job interface{}) {
		defer func() {
			recover()
		}()
		handler.Handle(job)
	}

	// The job worker.
	worker := func() {
		for {
			select {
			case job, ok := <-jobQueue:
				if !ok {
					return
				}
				handleJob(job)
			case <-cxt.Done():
				return
			}
		}
	}

	// Create the workers.
	for i := 0; i < workerNum; i++ {
		go worker()
	}
}
