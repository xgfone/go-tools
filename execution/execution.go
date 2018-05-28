// Package execution executes a command line program in a new process and returns a output.
package execution

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

// RunCmd executes the command, name, with its arguments, args,
// then returns stdout, stderr and error.
func RunCmd(cxt context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {

	cmd := exec.CommandContext(cxt, name, args...)
	var output bytes.Buffer
	var errput bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &errput
	err = cmd.Run()
	return output.Bytes(), errput.Bytes(), err
}

// Execute is the same as RunCmd, but only returns the error.
func Execute(cxt context.Context, name string, args ...string) error {
	_, stderr, err := RunCmd(cxt, name, args...)
	if err != nil {
		return fmt.Errorf("%s", string(stderr))
	}
	return nil
}

// Output is the same as RunCmd, but only returns the stdout and the error.
func Output(cxt context.Context, name string, args ...string) (string, error) {
	stdout, stderr, err := RunCmd(cxt, name, args...)
	if err != nil {
		return "", fmt.Errorf("%s", string(stderr))
	}
	return string(stdout), nil
}

// Executes is equal to Execute(cxt, cmds[0], cmds[1:]...)
func Executes(cxt context.Context, cmds []string) error {
	return Execute(cxt, cmds[0], cmds[1:]...)
}

// Outputs is equal to Output(cxt, cmds[0], cmds[1:]...).
func Outputs(cxt context.Context, cmds []string) (string, error) {
	return Output(cxt, cmds[0], cmds[1:]...)
}
