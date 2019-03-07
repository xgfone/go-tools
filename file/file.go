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

// Package file supplies some convenient functions about the file operation.
package file

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	// NotExist represents that the file or directory does not exist.
	NotExist = iota

	// FileType represents a file.
	FileType

	// DirType represents a directory.
	DirType
)

// HomeDir is the home directory of the current user.
var HomeDir = GetHomeDir()

// Type decides the type of a file.
//
// It returns FileType, DirType, or NotExist.
func Type(name string) uint8 {
	fi, err := os.Stat(name)
	if err == nil {
		if fi.IsDir() {
			return DirType
		}
		return FileType
	}
	if os.IsNotExist(err) {
		return NotExist
	}
	return FileType
}

// IsExist returns true if the file or directory exists, or return false.
func IsExist(filename string) bool {
	if Type(filename) == NotExist {
		return false
	}
	return true
}

// IsFile returns true if it's a file, or return false.
func IsFile(filename string) bool {
	if Type(filename) == FileType {
		return true
	}
	return false
}

// IsDir returns true if it's a directory, or return false.
func IsDir(filename string) bool {
	if Type(filename) == DirType {
		return true
	}
	return false
}

func addFile(lists []string, fullPath, fileName string, isfull bool) []string {
	if isfull {
		return append(lists, fullPath)
	}
	return append(lists, fileName)
}

// WalkDirFull returns all the filenames in a directory.
//
// dirPth is the directory where the file is in.
// If suffix is not empty, it only returns the files which have the suffix.
// If includeDir is true, it also returns the directory, not only the filename.
// If recursion is true, it will walk recursively.
// If fullPath is true, the filename is the full path, not only the name.
// If ignoreError is true, ignore the error; Or it will stop when an error occurs.
//
// Notice: the suffix is case-insensitive.
func WalkDirFull(dirPth, suffix string, includeDir, recursion, fullPath,
	ignoreError bool) ([]string, error) {

	files := make([]string, 0, 30)
	dirPth = strings.TrimRight(dirPth, "/")
	dirPth = strings.TrimRight(dirPth, "\\")
	if dirPth == "" {
		dirPth = "."
	}
	_, rootDir := filepath.Split(dirPth)

	suffix = strings.ToUpper(suffix)
	err := filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil && !ignoreError {
			return err
		}

		if fi.IsDir() {
			if fi.Name() == rootDir || recursion {
				return nil
			}

			if includeDir {
				files = addFile(files, filename, fi.Name(), fullPath)
			}

			return filepath.SkipDir
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = addFile(files, filename, fi.Name(), fullPath)
		}

		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}

// ListDir is the short for Walkdirfull, only recursion is false,
// fullpath is false, and ignoreerror is true.
func ListDir(dirPth, suffix string, includeDir bool) ([]string, error) {
	return WalkDirFull(dirPth, suffix, includeDir, false, false, true)
}

// ListDir2 is the short for Listdir, only suffix is empty, and includedir
// is false.
func ListDir2(dirPth string) ([]string, error) {
	return ListDir(dirPth, "", false)
}

// GetHomeDir returns the home directory.
//
// Return "" if the home direcotry is empty.
func GetHomeDir() string {
	if v := os.Getenv("HOME"); v != "" { // For Unix/Linux
		return v
	} else if v := os.Getenv("HOMEPATH"); v != "" { // For Windows
		return v
	}
	return ""
}

// Abs is similar to Abs in the std library "path/filepath",
// but firstly convert "~"" and "$HOME" to the home directory.
//
// Return the origin path if there is an error.
func Abs(p string) string {
	p = strings.TrimSpace(p)
	if p != "" && HomeDir != "" {
		_len := len(p)
		if p[0] == '~' {
			if _len == 1 || p[1] == '/' || p[1] == '\\' {
				p = filepath.Join(HomeDir, p[1:])
			}
		} else if _len >= 5 && p[:5] == "$HOME" {
			if _len == 5 || (p[5] == '/' || p[5] == '\\') {
				p = filepath.Join(HomeDir, p[5:])
			}
		}
	}

	if _p, err := filepath.Abs(p); err == nil {
		return _p
	}
	return p
}

// SelfPath returns the absolute path where the compiled executable file is in.
func SelfPath() string {
	return Abs(os.Args[0])
}

// SelfDir returns the directory where the compiled executable file is in.
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// EnsureDir make the directory if not exist.
func EnsureDir(fp string) error {
	return os.MkdirAll(fp, os.ModePerm)
}

// SearchFile searches a file in paths.
//
// Return nil if it didn't find the file.
//
// This is often used in search config file in /etc, ~/.
func SearchFile(filename string, paths ...string) []string {
	files := make([]string, 0, len(paths))
	for _, path := range paths {
		if fullPath := filepath.Join(path, filename); IsExist(fullPath) {
			files = append(files, fullPath)
		}
	}
	return files
}

// MTime returns the modified time of the file.
func MTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// Size returns the size of the file as how many bytes.
func Size(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}
