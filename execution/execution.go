// Package execution executes a command line program in a new process and returns a output.
package execution

import (
	"errors"
	"os/exec"
	"sync"
	"time"
)

var defaultGlobalMutex = new(sync.Mutex)

// ErrArguments is returned when the command failed to be executed.
var ErrArguments = errors.New("The arguments are too few")

// Execution executes a command line program.
type Execution struct {
	*sync.Mutex

	// Count stands for the times to be executed.
	//
	// The command will only be executed once whether or not success if it's 0.
	//
	// If it's a positive integer, the command will be executed repeatedly
	// until success or it has been executed by Count times.
	//
	// If it's a negative integer, the command will be executed repeatedly
	// until an error occurs or it has been executed by Count times.
	Count int

	// When retring, the it will be paused for that time.Millisecond.
	// If it's 0, don't pause.
	Interval int
}

// SetMutex replaces the sync.Mutex to a new one. The default is nil.
// When setting a mutex to non-nil, the execution will acquire the lock first
// before being executed.
//
// Set nil to clear the locker.
//
// Notice: This method is not thread-safe. Before replacing the mutex,
// ensave that the old one hasn't been locked.
func (e *Execution) SetMutex(m *sync.Mutex) {
	e.Mutex = m
}

// SetDefaultGlobalMutex is the same as e.SetMutex(defaultGlobalMutex).
func (e *Execution) SetDefaultGlobalMutex() {
	e.SetMutex(defaultGlobalMutex)
}

// Execute is the proxy of exec.Command(name, args...).Run(), but args[0] is name.
func (e *Execution) Execute(args []string) error {
	_, err := e.execute(args, false, func(name string, args []string, eout bool) (string, error) {
		err := exec.Command(name, args...).Run()
		return "", err
	})
	return err
}

// Output is the proxy of exec.Command(name, args...).Output(),
// but args[0] is name.
func (e *Execution) Output(args []string) (string, error) {
	return e.output(args, false)
}

// ErrOutput is the proxy of exec.Command(name, args...).CombinedOutput(),
// but args[0] is name.
func (e *Execution) ErrOutput(args []string) (string, error) {
	return e.output(args, true)
}

func (e *Execution) output(args []string, eout bool) (string, error) {
	return e.execute(args, eout, func(name string, args []string, eout bool) (string, error) {
		var err error
		var out []byte
		if eout {
			out, err = exec.Command(name, args...).CombinedOutput()
		} else {
			out, err = exec.Command(name, args...).Output()
		}
		return string(out), err
	})
}

func (e *Execution) execute(args []string, eout bool, executor func(string, []string, bool) (string, error)) (string, error) {
	if len(args) == 0 {
		return "", ErrArguments
	}
	name := args[0]
	args = args[1:]

	sleep := time.Millisecond * time.Duration(e.Interval)
	if sleep < 0 {
		sleep = 0
	}

	var count int
	var positive bool
	if e.Count < 0 {
		count = -e.Count
		positive = false
	} else {
		count = e.Count
		positive = true
	}

	var err error
	var out string

	for count >= 0 {
		func() {
			if e.Mutex != nil {
				e.Lock()
				defer e.Unlock()
			}
			out, err = executor(name, args, eout)
		}()

		if positive {
			if err == nil { // End until success
				break
			}
		} else {
			if err != nil { // End until failure
				break
			}
		}

		count--
		if count >= 0 && sleep > 0 {
			time.Sleep(sleep)
		}
	}

	return out, err
}
