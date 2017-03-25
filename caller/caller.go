// Package caller returns the filename and the line number where to call these functions.
package caller

import (
	"path"
	"runtime"
)

func getfileno(calldepth int) (string, int) {
	_, file, line, ok := runtime.Caller(calldepth + 1)
	if ok {
		return file, line
	}
	return "???", -1
}

// GetFileNo returns the filepath and the line number where calling it.
func GetFileNo() (string, int) {
	return getfileno(1)
}

// GetFile is the same as GetFileNo, but return the filepath.
func GetFile() string {
	file, _ := getfileno(1)
	return file
}

// GetNo is the same as GetFileNo, but return the line number.
func GetNo() int {
	_, line := getfileno(1)
	return line
}

// GetFileName is the same as GetFile, but return the filename, not the full path.
func GetFileName() string {
	file, _ := getfileno(1)
	_, filename := path.Split(file)
	return filename
}

// GetFileNameNo is the same as GetFileNo, but the filename, not the full path.
func GetFileNameNo() (string, int) {
	file, line := getfileno(1)
	_, filename := path.Split(file)
	return filename, line
}
