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

package file

import (
	"strings"
	"testing"

	"github.com/xgfone/go-tools/v7/slice"
)

func TestGetHomeDir(t *testing.T) {
	if home := GetHomeDir(); home == "" {
		t.Error("no home dir")
	}
}

func TestWalkDir(t *testing.T) {
	var err error
	var files []string

	// 1.1 => Test the match: suffix=""
	files, err = WalkDir("..", "", true, true, false, false)
	if err != nil {
		t.Error(err)
	} else if !slice.InStrings("file", files) {
		t.Error("not contain the directory 'file'")
	}

	// 1.2 => Test the match: suffix=".go"
	files, err = WalkDir("..", ".go", true, true, false, false)
	if err != nil {
		t.Error(err)
	} else if slice.InStrings("file", files) {
		t.Error("contain the directory 'file'")
	}

	// 2.1 => Test the match: fullPath=true
	files, err = WalkDir("..", "", true, true, true, false)
	if err != nil {
		t.Error(err)
	} else {
		for _, file := range files {
			if !strings.HasPrefix(file, "..") {
				t.Errorf("'%s' has no the directory prefix '%s'", file, "..")
			}
		}
	}

	// 2.2 => Test the match: fullPath=false
	files, err = WalkDir("..", "", true, true, false, false)
	if err != nil {
		t.Error(err)
	} else {
		for _, file := range files {
			if strings.HasPrefix(file, "..") {
				t.Errorf("'%s' has the directory prefix '%s'", file, "..")
			}
		}
	}

	// 3.1 => Test the match: includeDir=true, recursion=true
	files, err = WalkDir("..", "", true, true, false, false)
	if err != nil {
		t.Error(err)
	} else {
		if !slice.InStrings("file", files) {
			t.Error("not contain the directory 'file'")
		}
		if !slice.InStrings("file.go", files) {
			t.Error("not contain the directory 'file.go'")
		}
	}

	// 3.2 => Test the match: includeDir=true, recursion=false
	files, err = WalkDir("..", "", true, false, false, false)
	if err != nil {
		t.Error(err)
	} else {
		if !slice.InStrings("file", files) {
			t.Error("not contain the directory 'file'")
		}
		if slice.InStrings("file.go", files) {
			t.Error("contain the directory 'file.go'")
		}
	}

	// 3.3 => Test the match: includeDir=false, recursion=true
	files, err = WalkDir("..", "", false, true, false, false)
	if err != nil {
		t.Error(err)
	} else {
		if slice.InStrings("file", files) {
			t.Error("contain the directory 'file'")
		}
		if !slice.InStrings("file.go", files) {
			t.Error("not contain the directory 'file.go'")
		}
	}

	// 3.4 => Test the match: includeDir=false, recursion=false
	files, err = WalkDir("..", "", false, false, false, false)
	if err != nil {
		t.Error(err)
	} else {
		if slice.InStrings("file", files) {
			t.Error("contain the directory 'file'")
		}
		if slice.InStrings("file.go", files) {
			t.Error("contain the directory 'file.go'")
		}
	}

}
