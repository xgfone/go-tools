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
	return Execute(cxt, cmds[0], cmds[1:]...)
}

// Outputs is equal to Output(cxt, cmds[0], cmds[1:]...).
func Outputs(cxt context.Context, cmds []string) (string, error) {
	return Output(cxt, cmds[0], cmds[1:]...)
}
