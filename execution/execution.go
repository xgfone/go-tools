// execution executes a command line program in a new process and returns a output.
package execution

import (
	"errors"
	"os/exec"
	"time"
)

var ErrArguments = errors.New("The arguments are too few")

type Execution struct {
	// 0 stands for executing once only, and no retry if failed.
	Retry int

	// When retring, the it will be paused for that time.Millisecond.
	Interval int
}

func (e Execution) Execute(args []string) error {
	if len(args) == 0 {
		return ErrArguments
	}
	name := args[0]
	args = args[1:]

	sleep := time.Millisecond * time.Duration(e.Interval)
	retry := e.Retry
	ok := false
	var err error
	for retry >= 0 {
		if err = exec.Command(name, args...).Run(); err == nil {
			ok = true
			break
		}
		retry--
		if retry >= 0 {
			time.Sleep(sleep)
		}
	}

	if ok {
		return nil
	} else {
		return err
	}
}

func (e Execution) Output(args []string) (string, error) {
	return e.output(args, false)
}

func (e Execution) ErrOutput(args []string) (string, error) {
	return e.output(args, true)
}

func (e Execution) output(args []string, eout bool) (string, error) {
	if len(args) == 0 {
		return "", ErrArguments
	}
	name := args[0]
	args = args[1:]

	sleep := time.Millisecond * time.Duration(e.Interval)
	retry := e.Retry
	ok := false
	var err error
	var out []byte

	for retry >= 0 {
		if eout {
			out, err = exec.Command(name, args...).CombinedOutput()
		} else {
			out, err = exec.Command(name, args...).Output()
		}
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
		return string(out), nil
	} else {
		return "", err
	}
}
