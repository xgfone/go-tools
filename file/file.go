package file

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	NOT_EXIST = iota // Represent that the file does not exist.
	IS_FILE          // Represent a file.
	IS_DIR           // Represent a directory.
)

// Decide the type of a file.
//
// It returns IS_FILE, IS_DIR, or NOT_EXIST.
func FileType(name string) uint8 {
	fi, err := os.Stat(name)
	if err == nil {
		if fi.IsDir() {
			return IS_DIR
		}
		return IS_FILE
	}
	if os.IsNotExist(err) {
		return NOT_EXIST
	}
	return IS_FILE
}

// Return true if the file exists, or return false.
func IsExist(filename string) bool {
	if FileType(filename) == NOT_EXIST {
		return false
	}
	return true
}

// Return true if it's a file, or return false.
func IsFile(filename string) bool {
	if FileType(filename) == IS_FILE {
		return true
	}
	return false
}

// Return true if it's a directory, or return false.
func IsDir(filename string) bool {
	if FileType(filename) == IS_DIR {
		return true
	}
	return false
}

func addFile(lists []string, fullPath, fileName string, isfull bool) []string {
	if isfull {
		return append(lists, fullPath)
	} else {
		return append(lists, fileName)
	}
}

// Get the filename in a directory.
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

// The short for Walkdirfull, only recursion is false, fullpath is false,
// and ignoreerror is true.
func ListDir2(dirPth, suffix string, includeDir bool) ([]string, error) {
	return WalkDirFull(dirPth, suffix, includeDir, false, false, true)
}

// The short for Walkdirfull, only includedir is false, fullpath is true,
// and ignoreerror is true.
func WalkDir2(dirPth, suffix string, recursion bool) ([]string, error) {
	return WalkDirFull(dirPth, suffix, false, recursion, true, true)
}

// The short for Listdir2, only suffix is empty, and includedir is false.
func ListDir(dirPth string) ([]string, error) {
	return ListDir2(dirPth, "", false)
}

// The short for Walkdir2, only suffix is empty, and recursion is true.
func WalkDir(dirPth string) ([]string, error) {
	return WalkDir2(dirPth, "", true)
}

// Get the absolute path where the compiled executable file is in.
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// Get the absolute filepath, based on built executable file.
func RealPath(fp string) (string, error) {
	if path.IsAbs(fp) {
		return fp, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, fp), err
}

// Get the directory where the compiled executable file is in.
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// Mkdir dir if not exist
func EnsureDir(fp string) error {
	return os.MkdirAll(fp, os.ModePerm)
}

// Search a file in paths.
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

// Get the modified time of the file
func FileMTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// Get the size of the file as how many bytes
func FileSize(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}
