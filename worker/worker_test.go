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
