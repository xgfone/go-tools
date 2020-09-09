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

package execution

import "context"

// RunShellCmd is equal to DefaultCmd.RunShellCmd(ctx, name, args...).
func RunShellCmd(ctx context.Context, name string, args ...string) ([]byte, []byte, error) {
	return DefaultCmd.RunShellCmd(ctx, name, args...)
}

// RetryRunShellCmd is equal to DefaultCmd.RetryRunShellCmd(ctx, name, args...).
func RetryRunShellCmd(ctx context.Context, name string, args ...string) ([]byte, []byte, error) {
	return DefaultCmd.RetryRunShellCmd(ctx, name, args...)
}

// ExecuteShell is equal to DefaultCmd.ExecuteShell(ctx, name, args...).
func ExecuteShell(ctx context.Context, name string, args ...string) error {
	return DefaultCmd.ExecuteShell(ctx, name, args...)
}

// OutputShell is equal to DefaultCmd.OutputShell(ctx, name, args...).
func OutputShell(ctx context.Context, name string, args ...string) (string, error) {
	return DefaultCmd.OutputShell(ctx, name, args...)
}

// ExecutesShell is equal to DefaultCmd.ExecutesShell(ctx, cmds).
func ExecutesShell(ctx context.Context, cmds []string) error {
	return DefaultCmd.ExecutesShell(ctx, cmds)
}

// OutputsShell is equal to DefaultCmd.OutputsShell(ctx, cmds).
func OutputsShell(ctx context.Context, cmds []string) (string, error) {
	return DefaultCmd.OutputsShell(ctx, cmds)
}

// RunShellScript is equal to DefaultCmd.RunShellScript(ctx, name, args...).
func RunShellScript(ctx context.Context, name string, args ...string) ([]byte, []byte, error) {
	return DefaultCmd.RunShellScript(ctx, name, args...)
}

// RetryRunShellScript is equal to DefaultCmd.RetryRunShellScript(ctx, name, args...).
func RetryRunShellScript(ctx context.Context, name string, args ...string) ([]byte, []byte, error) {
	return DefaultCmd.RetryRunShellScript(ctx, name, args...)
}
