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

// Get the filepath and the line number where calling it.
func GetFileNo() (string, int) {
	return getfileno(1)
}

// Same as GetFileNo, but return the filepath.
func GetFile() string {
	file, _ := getfileno(1)
	return file
}

// Same as GetFileNo, but return the line number.
func GetNo() int {
	_, line := getfileno(1)
	return line
}

// Same as GetFile, but return the filename, not the full path.
func GetFileName() string {
	file, _ := getfileno(1)
	_, filename := path.Split(file)
	return filename
}

// Same as GetFileNo, but the filename, not the full path.
func GetFileNameNo() (string, int) {
	file, line := getfileno(1)
	_, filename := path.Split(file)
	return filename, line
}
