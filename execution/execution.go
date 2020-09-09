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
	"fmt"
	"os/exec"
	"sync"
	"time"
)

// ErrDeny is returned when the hook denies the cmd.
var ErrDeny = errors.New("the cmd is denied")

// CmdError represents a cmd error.
type CmdError struct {
	Name string
	Args []string

	Err    error
	Stdout []byte
	Stderr []byte
}

// NewCmdError returns a new CmdError.
func NewCmdError(name string, args []string, stdout, stderr []byte, err error) CmdError {
	return CmdError{Name: name, Args: args, Stdout: stdout, Stderr: stderr, Err: err}
}

func (c CmdError) Error() string {
	err := c.Err.Error()
	buf := bytes.NewBuffer(nil)
	buf.Grow(len(c.Stderr) + len(c.Stdout) + len(err) + len(c.Name) + 36)
	buf.WriteString("cmd=")
	buf.WriteString(c.Name)

	if len(c.Args) > 0 {
		fmt.Fprintf(buf, ", args=%s", c.Args)
	}
	if len(c.Stdout) > 0 {
		buf.WriteString(", stdout=")
		buf.Write(c.Stdout)
	}
	if len(c.Stderr) > 0 {
		buf.WriteString(", stderr=")
		buf.Write(c.Stderr)
	}

	buf.WriteString(", err=")
	buf.WriteString(err)
	return buf.String()
}

// Unwrap implements errors.Unwrap.
func (c CmdError) Unwrap() error {
	return c.Err
}

// Hook is used to filter or handle the cmd `name` with the arguments `args`.
//
// If returning true, it will continue to run it, or do nothing.
//
// DEPRECATED!!! Please use FilterHook.
type Hook func(name string, args ...string) bool

// FilterHook is used to filter or handle the cmd `name` with the arguments
// `args`, which returns the new name and args. If the returned name is empty,
// it will terminate to execute the cmd and do nothing.
type FilterHook func(name string, args ...string) (string, []string)

// ResultHook is used to filter the result and returns the new result.
type ResultHook func(name string, args []string, stdout, stderr []byte, err error) ([]byte, []byte, error)

// Cmd represents a command executor.
type Cmd struct {
	// DEPRECATED!!! Please use FilterHooks.
	Hooks []Hook

	FilterHooks []FilterHook
	ResultHooks []ResultHook

	// Shell is used to execute the command as the shell.
	//
	// If it's empty, it uses DefaultShell as the default.
	Shell string

	// Timeout is used to produce the timeout context based on the context
	// argument if not 0 when executing the command.
	Timeout time.Duration

	// SetCmd allows the user to customize exec.Cmd.
	//
	// Notice: You should not modify the fields `Stdout` and `Stderr`.
	SetCmd func(*exec.Cmd)

	// If Lock is not nil, it will be locked during the command is executed.
	Lock *sync.Mutex
}

// NewCmd returns a new executor Cmd.
func NewCmd() *Cmd {
	return new(Cmd)
}

// AppendHooks appends some hooks.
//
// DEPRECATED!!! Please use AppendFilterHooks.
func (c *Cmd) AppendHooks(hooks ...Hook) *Cmd {
	for _, hook := range hooks {
		if hook != nil {
			c.Hooks = append(c.Hooks, hook)
		}
	}
	return c
}

// AppendFilterHooks appends some filter hooks.
func (c *Cmd) AppendFilterHooks(hooks ...FilterHook) *Cmd {
	for _, hook := range hooks {
		if hook != nil {
			c.FilterHooks = append(c.FilterHooks, hook)
		}
	}
	return c
}

// AppendResultHooks appends some result hooks.
func (c *Cmd) AppendResultHooks(hooks ...ResultHook) *Cmd {
	for _, hook := range hooks {
		if hook != nil {
			c.ResultHooks = append(c.ResultHooks, hook)
		}
	}
	return c
}

func (c *Cmd) runCmd(cmd *exec.Cmd) error {
	if c.Lock != nil {
		c.Lock.Lock()
		defer c.Lock.Unlock()
	}
	return cmd.Run()
}

// RunCmd executes the command, name, with its arguments, args,
// then returns stdout, stderr and error.
//
// Notice: if there is an error to be returned, it is CmdError.
func (c *Cmd) RunCmd(cxt context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	if name == "" {
		panic("the cmd name is empty")
	}

	for _, hook := range c.Hooks {
		if ok := hook(name, args...); !ok {
			return c.runResultHooks(name, args, nil, nil, ErrDeny)
		}
	}

	for _, hook := range c.FilterHooks {
		_name, _args := hook(name, args...)
		if _name == "" {
			return c.runResultHooks(name, args, nil, nil, ErrDeny)
		}
		name, args = _name, _args
	}

	var cancel func()
	if c.Timeout > 0 {
		cxt, cancel = context.WithTimeout(cxt, c.Timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(cxt, name, args...)
	if c.SetCmd != nil {
		c.SetCmd(cmd)
	}

	var output bytes.Buffer
	var errput bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &errput
	err = c.runCmd(cmd)
	stdout = output.Bytes()
	stderr = errput.Bytes()

	return c.runResultHooks(name, args, stdout, stderr, err)
}

func (c *Cmd) runResultHooks(name string, args []string, stdout, stderr []byte,
	err error) ([]byte, []byte, error) {
	for _, hook := range c.ResultHooks {
		stdout, stderr, err = hook(name, args, stdout, stderr, err)
	}

	switch err.(type) {
	case nil, CmdError:
	default:
		err = NewCmdError(name, args, stdout, stderr, err)
	}

	return stdout, stderr, err
}

// Run is the alias of RunCmd.
func (c *Cmd) Run(ctx context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	return c.RunCmd(ctx, name, args...)
}

// RetryRunCmd is the same as RunCmd, but try to run once again if failed.
func (c *Cmd) RetryRunCmd(ctx context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	stdout, stderr, err = c.RunCmd(ctx, name, args...)
	if err != nil {
		stdout, stderr, err = c.RunCmd(ctx, name, args...)
	}
	return
}

// Execute is the same as RunCmd, but only returns the error.
func (c *Cmd) Execute(cxt context.Context, name string, args ...string) error {
	_, _, err := c.RunCmd(cxt, name, args...)
	return err
}

// Output is the same as RunCmd, but only returns the stdout and the error.
func (c *Cmd) Output(cxt context.Context, name string, args ...string) (string, error) {
	stdout, _, err := c.RunCmd(cxt, name, args...)
	return string(stdout), err
}

// Executes is equal to Execute(cxt, cmds[0], cmds[1:]...)
func (c *Cmd) Executes(cxt context.Context, cmds []string) error {
	_, _, err := c.RunCmd(cxt, cmds[0], cmds[1:]...)
	return err
}

// Outputs is equal to Output(cxt, cmds[0], cmds[1:]...).
func (c *Cmd) Outputs(cxt context.Context, cmds []string) (string, error) {
	stdout, _, err := c.RunCmd(cxt, cmds[0], cmds[1:]...)
	return string(stdout), err
}
