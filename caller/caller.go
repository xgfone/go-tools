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

func GetFileNo() (string, int) {
	return getfileno(1)
}

func GetFile() string {
	file, _ := getfileno(1)
	return file
}

func GetNo() int {
	_, line := getfileno(1)
	return line
}

func GetFileName() string {
	file, _ := getfileno(1)
	_, filename := path.Split(file)
	return filename
}

func GetFileNameNo() (string, int) {
	file, line := getfileno(1)
	_, filename := path.Split(file)
	return filename, line
}
