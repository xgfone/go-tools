package file

import (
	"os"
	"path"
)

// WriteBytes writes the byte content to a file.
//
// If successfully, return the byte nubmer to write and nil. Or return 0 and
// an error.
func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

// WriteString is the same as WriteBytes, but write the string to the file.
func WriteString(filePath string, s string) (int, error) {
	return WriteBytes(filePath, []byte(s))
}
