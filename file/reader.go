package file

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

// ToBytes reads a file by byte.
// error is not nil if there is a error.
func ToBytes(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// ToString reads a file by string.
// error is not nil if there is a error.
func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ToTrimString is the same as ToString, but remove the tail spaces.
func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}

// ToUint64 is the same as ToTrimString, but convert it to uint64.
func ToUint64(filePath string) (uint64, error) {
	content, err := ToTrimString(filePath)
	if err != nil {
		return 0, err
	}

	var ret uint64
	if ret, err = strconv.ParseUint(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}

// ToInt64 is the same as ToTrimString, but convert it to int64.
func ToInt64(filePath string) (int64, error) {
	content, err := ToTrimString(filePath)
	if err != nil {
		return 0, err
	}

	var ret int64
	if ret, err = strconv.ParseInt(content, 10, 64); err != nil {
		return 0, err
	}
	return ret, nil
}

// ReadLine reads the content in the buffer by line.
func ReadLine(r *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := r.ReadLine()
	for isPrefix && err == nil {
		var bs []byte
		bs, isPrefix, err = r.ReadLine()
		line = append(line, bs...)
	}

	return line, err
}
