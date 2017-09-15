package worker

import (
	"context"
	"fmt"
	"time"
)

func ExampleDispatch() {
	cxt, cancel := context.WithCancel(context.TODO())
	jobQueue := make(chan interface{}, 10)
	Dispatch(cxt, 2, jobQueue, TaskFunc(func(job interface{}) {
		fmt.Println(job)
	}))

	jobQueue <- 11
	time.Sleep(time.Millisecond * 10)
	jobQueue <- "aa"
	time.Sleep(time.Millisecond * 10)
	cancel()

	// Output:
	// 11
	// aa
}

func ExampleDispatcher() {
	JobQueue := make(chan interface{}, 2)
	dispatcher := NewDispatcher(5, TaskFunc(func(job interface{}) {
		//fmt.Printf("Receive job: %v\n", job.Payload)
		fmt.Printf("Receive job\n")
	}))
	dispatcher.Dispatch(JobQueue)

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		JobQueue <- i
	}

	time.Sleep(time.Second)
	dispatcher.Stop()

	// Output:
	// Receive job
	// Receive job
	// Receive job
	// Receive job
	// Receive job
	// Receive job
	// Receive job
	// Receive job
	// Receive job
	// Receive job
}
