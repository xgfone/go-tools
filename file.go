package utils

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
