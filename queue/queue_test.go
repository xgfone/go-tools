package queue

import (
	"fmt"
)

func ExampleNewMemoryQueue() {
	q := NewMemoryQueue(0)
	exit := make(chan struct{})
	go func() {
		for {
			v, err := q.Get()
			if err != nil || v == nil {
				fmt.Println("exit")
				exit <- struct{}{}
				return
			}
			fmt.Println(v)
		}
	}()

	q.Put("abc")
	q.Put(123)
	q.Put(nil)
	<-exit

	// Output:
	// abc
	// 123
	// exit
}
