package sync2

import (
	"fmt"
	"sync"
	"time"
)

func ExampleResourceLock_Lock() {
	locker := NewResourceLock()
	id := "ID0123456789"
	waiter := sync.WaitGroup{}

	waiter.Add(1)
	go func() {
		locker.Lock(id)
		fmt.Println("Locked")
		time.Sleep(time.Second)
		fmt.Println("Unlocking")
		locker.Unlock(id)
		waiter.Done()
	}()

	waiter.Add(1)
	go func() {
		locker.Lock(id)
		fmt.Println("Locked")
		time.Sleep(time.Second)
		fmt.Println("Unlocking")
		locker.Unlock(id)
		waiter.Done()
	}()

	waiter.Add(1)
	go func() {
		locker.Lock(id)
		fmt.Println("Locked")
		time.Sleep(time.Second)
		fmt.Println("Unlocking")
		locker.Unlock(id)
		waiter.Done()
	}()

	waiter.Wait()

	// Output:
	// Locked
	// Unlocking
	// Locked
	// Unlocking
	// Locked
	// Unlocking
}
