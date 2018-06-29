// Package handler is the handler collection of the logger.
package handler

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/xgfone/go-tools/file"
	"github.com/xgfone/go-tools/function"
)

// FileMode is the mode to open the log file.
const FileMode = os.O_APPEND | os.O_CREATE | os.O_WRONLY

// FilePerm is the default permission to open the log file.
var FilePerm os.FileMode = 0644

// ErrFileNotOpen is the error to open the log file.
var ErrFileNotOpen = errors.New("The file is not opened")

// SizedRotatingFile is a rotating logging handler based on the size.
type SizedRotatingFile struct {
	sync.Mutex
	w *WriteCloser

	filename    string
	maxSize     int
	backupCount int
	nbytes      int
	buffer      bool
}

// NewSizedRotatingFile returns a new RotatingFile.
func NewSizedRotatingFile(filename string, size, count int) *SizedRotatingFile {
	r := &SizedRotatingFile{
		filename:    filename,
		maxSize:     size,
		backupCount: count,
		buffer:      true,
	}

	if err := r.open(); err != nil {
		panic(err)
	}
	return r
}

// DisableBuffer disables the buffer, which writes the log into the file
// immediately.
func (r *SizedRotatingFile) DisableBuffer() {
	r.Lock()
	r.buffer = false
	r.Unlock()
}

// Write implements the interface io.Writer.
func (r *SizedRotatingFile) Write(data []byte) (n int, err error) {
	return r.WriteString(string(data))
}

// WriteString writes the string.
func (r *SizedRotatingFile) WriteString(data string) (n int, err error) {
	r.Lock()
	defer r.Unlock()

	if r.w == nil || r.w.Closed() {
		err = ErrFileNotOpen
		return
	}

	if r.nbytes+len(data) > r.maxSize {
		if err = r.doRollover(); err != nil {
			return
		}
	}

	if n, err = r.w.WriteString(data); err != nil {
		return
	}
	if !r.buffer {
		r.w.Flush()
	}
	r.nbytes += n
	return
}

// Close implements the interface io.Closer.
func (r *SizedRotatingFile) Close() (err error) {
	r.Lock()
	err = r.close()
	r.Unlock()
	return
}

// Flush flushes the log cache into the file.
func (r *SizedRotatingFile) Flush() (err error) {
	r.Lock()
	err = r.w.Flush()
	r.Unlock()
	return
}

// ReOpen reopens the log file.
//
// Notice: If it failed to reopen the log file, the log will be redirected to
// os.Stderr.
func (r *SizedRotatingFile) ReOpen() (err error) {
	r.Lock()
	r.close()
	err = r.open()
	r.Unlock()
	return
}

func (r *SizedRotatingFile) close() (err error) {
	if r.w != nil {
		err = r.w.Close()
		r.w = nil
	}
	return
}

func (r *SizedRotatingFile) doRollover() (err error) {
	if r.backupCount > 0 {
		if err = r.close(); err != nil {
			return fmt.Errorf("Rotating: close failed: %s", err)
		}

		if !file.IsExist(r.filename) {
			return nil
		} else if n, err := file.Size(r.filename); err != nil {
			return fmt.Errorf("Rotating: failed to get the size: %s", err)
		} else if n == 0 {
			return nil
		}

		for _, i := range function.Range(r.backupCount-1, 0, -1) {
			sfn := fmt.Sprintf("%s.%d", r.filename, i)
			dfn := fmt.Sprintf("%s.%d", r.filename, i+1)
			if file.IsExist(sfn) {
				if file.IsExist(dfn) {
					os.Remove(dfn)
				}
				if err = os.Rename(sfn, dfn); err != nil {
					return fmt.Errorf("Rotating: failed to rename %s -> %s: %s",
						sfn, dfn, err)
				}
			}
		}
		dfn := r.filename + ".1"
		if file.IsExist(dfn) {
			if err = os.Remove(dfn); err != nil {
				return fmt.Errorf("Rotating: failed to remove %s: %s", dfn, err)
			}
		}
		if file.IsExist(r.filename) {
			if err = os.Rename(r.filename, dfn); err != nil {
				return fmt.Errorf("Rotating: failed to rename %s -> %s: %s",
					r.filename, dfn, err)
			}
		}
		err = r.open()
	}
	return
}

func (r *SizedRotatingFile) open() (err error) {
	file, err := os.OpenFile(r.filename, FileMode, FilePerm)
	if err != nil {
		return
	}
	info, err := file.Stat()
	if err != nil {
		return
	}
	r.nbytes = int(info.Size())
	r.w = NewWriteCloser(file)
	return
}
