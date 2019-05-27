// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package execution executes a command line program in a new process and returns a output.
package execution

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
)

// ErrDeny is returned when the hook denies the cmd.
var ErrDeny = errors.New("the cmd is denied")

var globalHooks []Hook

// Hook is used to filter or handle the cmd `name` with the arguments `args`.
//
// If returning true, it will continue to run it, or do nothing.
type Hook func(name string, args ...string) bool

// AppendHook appends the hooks to the global hooks.
func AppendHook(hooks ...Hook) {
	for _, hook := range hooks {
		if hook != nil {
			globalHooks = append(globalHooks, hook)
		}
	}
}

func geterr(stdout, stderr []byte, err error) error {
	if err != nil {
		if len(stderr) > 0 {
			err = errors.New(string(stderr))
		} else if len(stdout) > 0 {
			err = errors.New(string(stdout))
		}
	}
	return err
}

// RunCmd executes the command, name, with its arguments, args,
// then returns stdout, stderr and error.
func RunCmd(cxt context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {

	for _, hook := range globalHooks {
		if ok := hook(name, args...); !ok {
			return nil, nil, ErrDeny
		}
	}

	cmd := exec.CommandContext(cxt, name, args...)
	var output bytes.Buffer
	var errput bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &errput
	err = cmd.Run()
	stdout = output.Bytes()
	stderr = errput.Bytes()
	err = geterr(stdout, stderr, err)
	return
}

// Run is the alias of RunCmd.
func Run(ctx context.Context, name string, args ...string) ([]byte, []byte, error) {
	return RunCmd(ctx, name, args...)
}

// RetryRunCmd is the same as RunCmd, but try to run once again if failed.
func RetryRunCmd(ctx context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	stdout, stderr, err = RunCmd(ctx, name, args...)
	if err != nil {
		stdout, stderr, err = RunCmd(ctx, name, args...)
	}
	return
}

// Execute is the same as RunCmd, but only returns the error.
func Execute(cxt context.Context, name string, args ...string) error {
	_, _, err := RunCmd(cxt, name, args...)
	return err
}

// Output is the same as RunCmd, but only returns the stdout and the error.
func Output(cxt context.Context, name string, args ...string) (string, error) {
	stdout, _, err := RunCmd(cxt, name, args...)
	return string(stdout), err
}

// Executes is equal to Execute(cxt, cmds[0], cmds[1:]...)
func Executes(cxt context.Context, cmds []string) error {
	_, _, err := RunCmd(cxt, cmds[0], cmds[1:]...)
	return err
}

// Outputs is equal to Output(cxt, cmds[0], cmds[1:]...).
func Outputs(cxt context.Context, cmds []string) (string, error) {
	stdout, _, err := RunCmd(cxt, cmds[0], cmds[1:]...)
	return string(stdout), err
}
