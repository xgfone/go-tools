// Package handler is the handler collection of the logger.
package handler

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"sync"
	"time"

	"github.com/xgfone/go-tools/file"
	"github.com/xgfone/go-tools/function"
)

const (
	// FileMode is the mode to open the log file.
	FileMode = os.O_APPEND | os.O_CREATE | os.O_WRONLY

	// FilePerm is the default permission to open the log file.
	FilePerm os.FileMode = 0644
)

var (
	dayRE       = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(\.\w+)?$`)
	day   int64 = 3600 * 24

	time2fmt = map[int64]string{
		day: "2006-01-02",
	}

	filePerm = FilePerm
)

var (
	// ErrFileNotOpen is the error to open the log file.
	ErrFileNotOpen = errors.New("The file is not opened")
)

// ResetDefaultFilePerm resets the default permission to open the log file.
func ResetDefaultFilePerm(perm int) {
	filePerm = os.FileMode(perm)
}

// TimedRotatingFile is a file handler based on the timed rotating, like
// `logging.handlers.TimedRotatingFileHandler` in Python.
// Now only support the rotation by day.
//
// DEPRECATED!!! Not suggest to use it again. Use SizedRotatingFile instead.
type TimedRotatingFile struct {
	sync.Mutex
	w io.WriteCloser

	filename    string
	backupCount int
	interval    int64
	when        int64
	rotatorAt   int64
	extRE       *regexp.Regexp
}

// NewTimedRotatingFile creates a new TimedRotatingFile.
//
// If failed, it will panic.
func NewTimedRotatingFile(filename string, count int) *TimedRotatingFile {
	filename = file.Abs(filename)
	t := TimedRotatingFile{
		filename:    filename,
		when:        day,
		extRE:       dayRE,
		backupCount: count,
		interval:    day,
	}
	t.reComputeRollover()
	if err := t.open(); err != nil {
		panic(err)
	}
	return &t
}

// WriteString writes the string data into the file, which may rotate the file if necessary.
func (t *TimedRotatingFile) WriteString(data string) (n int, err error) {
	return t.Write([]byte(data))
}

// Write writes the byte slice data into the file, which may rotate the file if necessary.
func (t *TimedRotatingFile) Write(data []byte) (n int, err error) {
	t.Lock()
	defer t.Unlock()

	if t.w == nil {
		err = ErrFileNotOpen
		return
	}

	if t.shouldRollover() {
		if err = t.doRollover(); err != nil {
			return
		}
	}

	return t.w.Write(data)
}

func (t *TimedRotatingFile) shouldRollover() bool {
	return time.Now().Unix() >= t.rotatorAt
}

// Close closes the handler.
// Return ErrFileNotOpen when to write the data to the handler after closed.
func (t *TimedRotatingFile) Close() (err error) {
	if err = t.w.Close(); err != nil {
		return
	}
	t.w = nil
	return
}

func (t *TimedRotatingFile) open() error {
	file, err := os.OpenFile(t.filename, FileMode, FilePerm)
	if err != nil {
		return err
	}
	t.w = file
	return nil
}

func (t *TimedRotatingFile) doRollover() (err error) {
	if err = t.Close(); err != nil {
		return
	}

	dstTime := t.rotatorAt - t.interval
	dstPath := t.filename + "." + time.Unix(dstTime, 0).Format(time2fmt[t.when])
	if file.IsExist(dstPath) {
		os.Remove(dstPath)
	}

	if file.IsFile(t.filename) {
		if err = os.Rename(t.filename, dstPath); err != nil {
			return err
		}
	}

	if t.backupCount > 0 {
		for _, file := range t.getFilesToDelete() {
			os.Remove(file)
		}
	}

	t.reComputeRollover()
	return t.open()
}

func (t *TimedRotatingFile) getFilesToDelete() []string {
	result := make([]string, 0, 30)
	dirName, baseName := filepath.Split(t.filename)
	fileNames, err := file.ListDir2(dirName)
	if err != nil {
		return result
	}

	var suffix, prefix string
	_prefix := baseName + "."
	plen := len(_prefix)
	for _, fileName := range fileNames {
		if len(fileName) <= plen {
			continue
		}
		prefix = string(fileName[:plen])
		if _prefix == prefix {
			suffix = string(fileName[plen:])
			if t.extRE.MatchString(suffix) {
				result = append(result, filepath.Join(dirName, fileName))
			}
		}
	}

	if len(result) <= t.backupCount {
		return []string{}
	}
	sort.Strings(result)
	return result[:len(result)-t.backupCount]
}

func (t *TimedRotatingFile) reComputeRollover() {
	currentTime := time.Now().Unix()

	_time := time.Unix(currentTime, 0)
	currentHour := _time.Hour()
	currentMinute := _time.Minute()
	currentSecond := _time.Second()

	r := t.interval - int64((currentHour*60+currentMinute)*60+currentSecond)
	t.rotatorAt = currentTime + r
}

// SizedRotatingFile is a rotating logging handler based on the size.
type SizedRotatingFile struct {
	sync.Mutex
	w *WriteCloser

	filename    string
	maxSize     int
	backupCount int
	nbytes      int
}

// NewSizedRotatingFile returns a new RotatingFile.
func NewSizedRotatingFile(filename string, size, count int) *SizedRotatingFile {
	r := &SizedRotatingFile{
		filename:    filename,
		maxSize:     size,
		backupCount: count,
	}

	if err := r.open(); err != nil {
		panic(err)
	}
	return r
}

// Write implements the interface io.Writer.
func (r *SizedRotatingFile) Write(data []byte) (n int, err error) {
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

	if n, err = r.w.Write(data); err != nil {
		return
	}
	r.nbytes += n
	return
}

// WriteString writes the string.
func (r *SizedRotatingFile) WriteString(data string) (n int, err error) {
	return r.Write([]byte(data))
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
			return
		}
		for _, i := range function.Range(r.backupCount-1, 0, -1) {
			sfn := fmt.Sprintf("%s.%d", r.filename, i)
			dfn := fmt.Sprintf("%s.%d", r.filename, i+1)
			if file.IsExist(sfn) {
				if file.IsExist(dfn) {
					os.Remove(dfn)
				}
				if err = os.Rename(sfn, dfn); err != nil {
					return
				}
			}
		}
		dfn := r.filename + ".1"
		if file.IsExist(dfn) {
			if err = os.Remove(dfn); err != nil {
				return
			}
		}
		if file.IsExist(r.filename) {
			if err = os.Rename(r.filename, dfn); err != nil {
				return
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
