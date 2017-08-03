// Package file supplies some convenient functions about the file operation.
package file

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	// NotExist represents that the file does not exist.
	NotExist = iota

	// IsFileType represents a file.
	IsFileType

	// IsDirType represents a directory.
	IsDirType
)

// HomeDir is the home directory of the current user.
var HomeDir = GetHomeDir()

// Type decides the type of a file.
//
// It returns IsFileType, IsDirType, or NotExist.
func Type(name string) uint8 {
	fi, err := os.Stat(name)
	if err == nil {
		if fi.IsDir() {
			return IsDirType
		}
		return IsFileType
	}
	if os.IsNotExist(err) {
		return NotExist
	}
	return IsFileType
}

// IsExist returns true if the file exists, or return false.
func IsExist(filename string) bool {
	if Type(filename) == NotExist {
		return false
	}
	return true
}

// IsFile returns true if it's a file, or return false.
func IsFile(filename string) bool {
	if Type(filename) == IsFileType {
		return true
	}
	return false
}

// IsDir returns true if it's a directory, or return false.
func IsDir(filename string) bool {
	if Type(filename) == IsDirType {
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
// If suffix is not empty, it only returns the files which have the suffix of 'suffix'.
// If includedir is true, it also returns the directory, not only the filename.
// If recursion is true, it will walk recursively.
// If fullpath is true, the filename is the full path, not only the name.
// If ignoreerror is true, ignore the error; Or it will stop when a error occurs.
func WalkDirFull(dirPth, suffix string, includeDir, recursion, fullPath, ignoreError bool) ([]string, error) {
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

// ListDir2 is the short for Walkdirfull, only recursion is false, fullpath is false,
// and ignoreerror is true.
func ListDir2(dirPth, suffix string, includeDir bool) ([]string, error) {
	return WalkDirFull(dirPth, suffix, includeDir, false, false, true)
}

// WalkDir2 is the short for Walkdirfull, only includedir is false, fullpath is true,
// and ignoreerror is true.
func WalkDir2(dirPth, suffix string, recursion bool) ([]string, error) {
	return WalkDirFull(dirPth, suffix, false, recursion, true, true)
}

// ListDir is the short for Listdir2, only suffix is empty, and includedir is false.
func ListDir(dirPth string) ([]string, error) {
	return ListDir2(dirPth, "", false)
}

// WalkDir is the short for Walkdir2, only suffix is empty, and recursion is true.
func WalkDir(dirPth string) ([]string, error) {
	return WalkDir2(dirPth, "", true)
}

// GetHomeDir returns the home directory. Return "" if the home direcotry is empty.
func GetHomeDir() string {
	if v := os.Getenv("HOME"); v != "" { // For Unix/Linux
		return v
	} else if v := os.Getenv("HOMEPATH"); v != "" { // For Windows
		return v
	}
	return ""
}

// Abs is similar to Abs in the std library "path/filepath", but firstly convert
// "~"" and "$HOME" to the home directory. Return the origin path if there is an
// error.
func Abs(p string) string {
	if HomeDir != "" {
		p = strings.Replace(p, "~", HomeDir, 1)
		p = strings.Replace(p, "$HOME", HomeDir, 1)
	}

	if _p, err := filepath.Abs(p); err == nil {
		return _p
	}
	return p
}

// SelfPath returns the absolute path where the compiled executable file is in.
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// RealPath returns the absolute filepath, based on built executable file.
func RealPath(fp string) (string, error) {
	if path.IsAbs(fp) {
		return fp, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, fp), err
}

// SelfDir returns the directory where the compiled executable file is in.
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// EnsureDir make the directory if not exist
func EnsureDir(fp string) error {
	return os.MkdirAll(fp, os.ModePerm)
}

// SearchFile searches a file in paths.
// This is often used in search config file in /etc, ~/.
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = fmt.Errorf("%s not found in paths", fullPath)
	return
}

// MTime returns the modified time of the file
func MTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// Size returns the size of the file as how many bytes
func Size(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}
