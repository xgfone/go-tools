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

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// DefaultShell is the default shell to execute the shell command or script.
var DefaultShell = "bash"

// ShellScriptDir is the directory to save the script file to be executed.
//
// If OS is windows or js, it is reset to "". But you can set it to somewhere.
var ShellScriptDir = os.TempDir()

func init() {
	switch runtime.GOOS {
	case "js", "windows":
		ShellScriptDir = ""
	}
}

// RunShellCmd runs the command name with args as the shell command, that's,
//   shell -c "fmt.Sprintf(name, args...)".
func (c *Cmd) RunShellCmd(ctx context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	shell := c.Shell
	if shell == "" {
		shell = DefaultShell
	}

	if _len := len(args); _len != 0 {
		vs := make([]interface{}, _len)
		for i := 0; i < _len; i++ {
			vs[i] = args[i]
		}
		name = fmt.Sprintf(name, vs...)
	}

	return c.RunCmd(ctx, shell, "-c", name)
}

// RetryRunShellCmd is the same as RunShellCmd, but try to run once again if failed.
func (c *Cmd) RetryRunShellCmd(ctx context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	if stdout, stderr, err = c.RunShellCmd(ctx, name, args...); err != nil {
		stdout, stderr, err = c.RunShellCmd(ctx, name, args...)
	}
	return
}

// ExecuteShell is the same as RunShellCmd, but only returns the error.
func (c *Cmd) ExecuteShell(cxt context.Context, name string, args ...string) error {
	_, _, err := c.RunShellCmd(cxt, name, args...)
	return err
}

// OutputShell is the same as RunShellCmd, but only returns the stdout and the error.
func (c *Cmd) OutputShell(cxt context.Context, name string, args ...string) (string, error) {
	stdout, _, err := c.RunShellCmd(cxt, name, args...)
	return string(stdout), err
}

// ExecutesShell is equal to ExecuteShell(cxt, cmds[0], cmds[1:]...)
func (c *Cmd) ExecutesShell(cxt context.Context, cmds []string) error {
	return c.ExecuteShell(cxt, cmds[0], cmds[1:]...)
}

// OutputsShell is equal to OutputShell(cxt, cmds[0], cmds[1:]...).
func (c *Cmd) OutputsShell(cxt context.Context, cmds []string) (string, error) {
	return c.OutputShell(cxt, cmds[0], cmds[1:]...)
}

// RetryRunShellScript is the same as RunShellScript, but try to run once again
// if failed.
func (c *Cmd) RetryRunShellScript(ctx context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	if stdout, stderr, err = c.RunShellScript(ctx, name, args...); err != nil {
		stdout, stderr, err = c.RunShellScript(ctx, name, args...)
	}
	return
}

// RunShellScript runs the command name with args as the shell script,
// the content of which is fmt.Sprintf(name, args...).
func (c *Cmd) RunShellScript(ctx context.Context, name string, args ...string) (
	stdout, stderr []byte, err error) {
	script := name
	if _len := len(args); _len != 0 {
		vs := make([]interface{}, _len)
		for i := 0; i < _len; i++ {
			vs[i] = args[i]
		}
		script = fmt.Sprintf(name, vs...)
	}

	filename, err := c.getScriptFile(script)
	if err != nil {
		err = NewCmdError(name, args, nil, nil, err)
		return
	}
	defer os.RemoveAll(filename)

	shell := c.Shell
	if shell == "" {
		shell = DefaultShell
	}
	return c.RunCmd(ctx, shell, filename)
}

func (c *Cmd) getScriptFile(script string) (filename string, err error) {
	data := []byte(script)
	md5sum := md5.Sum(data)
	hexsum := hex.EncodeToString(md5sum[:])
	filename = fmt.Sprintf("__execution_run_shell_script_%s.sh", hexsum)

	if ShellScriptDir != "" {
		filename = filepath.Join(ShellScriptDir, filename)
	}

	err = ioutil.WriteFile(filename, data, 0700)
	return
}
