// Copyright 2020 xgfone
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

package namespace

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xgfone/go-tools/v7/execution"
)

// ErrNoNameSpace is returned when no namespace associates with a given process.
var ErrNoNameSpace = errors.New("no namespace")

func execute(name string, args ...string) (string, error) {
	return execution.Output(context.Background(), name, args...)
}

// NameSpace is the Linux NameSpace instance.
type NameSpace struct {
	Name string
}

// NewNameSpace returns a new NameSpace instance named name.
func NewNameSpace(name string) NameSpace { return NameSpace{Name: name} }

// NewNameSpaceFromPid returns a new NameSpace instance by the process pid.
func NewNameSpaceFromPid(pid int) (NameSpace, error) {
	out, err := execute("ip", "netns", "identify", fmt.Sprint(pid))
	if err != nil {
		return NameSpace{}, err
	}

	if out = strings.TrimSpace(out); out == "" {
		return NameSpace{}, ErrNoNameSpace
	}
	return NameSpace{Name: out}, nil
}

// GetAllNameSpace returns all the namespace instances.
func GetAllNameSpace() (nss []NameSpace, err error) {
	nss = make([]NameSpace, 0, 128)
	err = filepath.Walk("/var/run/netns", func(_ string, info os.FileInfo, err error) error {
		nss = append(nss, NameSpace{Name: info.Name()})
		return err
	})
	return
}

func (ns NameSpace) String() string {
	return fmt.Sprintf("NameSpace(%s)", ns.Name)
}

// Pids returns all the pids in the current namespace.
func (ns NameSpace) Pids() (pids []int, err error) {
	out, err := execute("ip", "netns", "pids", ns.Name)
	if err == nil {
		var v int64
		ss := strings.Fields(out)
		pids = make([]int, len(ss))
		for i, s := range ss {
			if v, err = strconv.ParseInt(s, 10, 64); err != nil {
				return
			}
			pids[i] = int(v)
		}
	}
	return
}

// Create creates the namespace.
func (ns NameSpace) Create() (err error) {
	if _, err = execute("ip", "netns", "add", ns.Name); err != nil {
		if strings.Contains(err.Error(), "File exists") {
			_, err = ns.Exec("ip", "link", "set", "lo", "up")
		}
	}
	return
}

// Delete deletes the namespance.
func (ns NameSpace) Delete() (err error) {
	if _, err = execute("ip", "netns", "delete", ns.Name); err != nil {
		if strings.Contains(err.Error(), "No such file or directory") {
			err = nil
		}
	}
	return
}

// IsExist reports whether the namespace exists or not.
func (ns NameSpace) IsExist() (exist bool, err error) {
	if _, err = os.Stat("/var/run/netns/" + ns.Name); err == nil {
		exist = true
	} else if os.IsNotExist(err) {
		err = nil
	}

	return
}

// Exec executes a shell command in the current namespace.
func (ns NameSpace) Exec(cmd string, args ...string) (output string, err error) {
	_args := make([]string, 4+len(args))
	_args[0] = "netns"
	_args[1] = "exec"
	_args[2] = ns.Name
	_args[3] = cmd
	copy(_args[4:], args)
	return execute("ip", _args...)
}
