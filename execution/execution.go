// execution executes a command line program in a new process and returns a output.
package execution

import (
	"errors"
	"os/exec"
	"sync"
	"time"
)

var ErrArguments = errors.New("The arguments are too few")

type Execution struct {
	*sync.Mutex

	// 0 stands for executing once only, and no retry if failed.
	Retry int

	// When retring, the it will be paused for that time.Millisecond.
	Interval int

	// Whether acquire the lock firstly when executing.
	IsLock bool
}

func (e *Execution) Execute(args []string) error {
	_, err := e.execute(args, false, func(name string, args []string, eout bool) (string, error) {
		err := exec.Command(name, args...).Run()
		return "", err
	})
	return err
}

func (e *Execution) Output(args []string) (string, error) {
	return e.output(args, false)
}

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

	retry := e.Retry
	ok := false
	var err error
	var out string

	for retry >= 0 {
		func() {
			if e.IsLock {
				e.Lock()
				defer e.Unlock()
			}
			out, err = executor(name, args, eout)
		}()

		if err == nil {
			ok = true
			break
		}

		retry--
		if retry >= 0 {
			time.Sleep(sleep)
		}
	}

	if ok {
		return out, nil
	} else {
		return "", err
	}
}
