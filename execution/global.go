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

import "context"

// DefaultCmd is the global default cmd executor.
var DefaultCmd = NewCmd()

// AppendHooks is equal to DefaultCmd.AppendHook(hooks...).
func AppendHooks(hooks ...Hook) {
	DefaultCmd.AppendHooks(hooks...)
}

// RunCmd is equal to DefaultCmd.RunCmd(ctx, name, args...).
func RunCmd(ctx context.Context, name string, args ...string) (stdout, stderr []byte, err error) {
	return DefaultCmd.RunCmd(ctx, name, args...)
}

// Run is the alias of RunCmd.
func Run(ctx context.Context, name string, args ...string) (stdout, stderr []byte, err error) {
	return DefaultCmd.Run(ctx, name, args...)
}

// RetryRunCmd is equal to DefaultCmd.RetryRunCmd(ctx, name, args...).
func RetryRunCmd(ctx context.Context, name string, args ...string) (stdout, stderr []byte, err error) {
	return DefaultCmd.RetryRunCmd(ctx, name, args...)
}

// Execute is equal to DefaultCmd.Execute(cxt, name, args...).
func Execute(cxt context.Context, name string, args ...string) error {
	return DefaultCmd.Execute(cxt, name, args...)
}

// Output is equal to DefaultCmd.Output(cxt, name, args...).
func Output(cxt context.Context, name string, args ...string) (string, error) {
	return DefaultCmd.Output(cxt, name, args...)
}

// Executes is equal to DefaultCmd.Executes(cxt, cmds).
func Executes(cxt context.Context, cmds []string) error {
	return DefaultCmd.Executes(cxt, cmds)
}

// Outputs is equal to DefaultCmd.Outputs(cxt, cmds).
func Outputs(cxt context.Context, cmds []string) (string, error) {
	return DefaultCmd.Outputs(cxt, cmds)
}
