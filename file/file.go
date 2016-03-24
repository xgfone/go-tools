package file

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	NOT_EXIST = iota
	IS_FILE
	IS_DIR
)

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

func IsExist(filename string) bool {
	if FileType(filename) == NOT_EXIST {
		return false
	}
	return true
}

func IsFile(filename string) bool {
	if FileType(filename) == IS_FILE {
		return true
	}
	return false
}

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

func ListDir2(dirPth, suffix string, includeDir bool) ([]string, error) {
	return WalkDirFull(dirPth, suffix, includeDir, false, false, true)
}

func WalkDir2(dirPth, suffix string, recursion bool) ([]string, error) {
	return WalkDirFull(dirPth, suffix, false, recursion, true, true)
}

func ListDir(dirPth string) ([]string, error) {
	return ListDir2(dirPth, "", false)
}

func WalkDir(dirPth string) ([]string, error) {
	return WalkDir2(dirPth, "", true)
}

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// get absolute filepath, based on built executable file
func RealPath(fp string) (string, error) {
	if path.IsAbs(fp) {
		return fp, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, fp), err
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// mkdir dir if not exist
func EnsureDir(fp string) error {
	return os.MkdirAll(fp, os.ModePerm)
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = fmt.Errorf("%s not found in paths", fullPath)
	return
}

// get file modified time
func FileMTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}
