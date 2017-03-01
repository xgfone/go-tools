package worker_test

import (
	"fmt"
	"time"

	"github.com/xgfone/go-tools/worker"
)

func ExampleDispatcher() {
	JobQueue := make(chan interface{}, 2)
	dispatcher := worker.NewDispatcher(5, worker.FuncTask(func(job interface{}) {
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
