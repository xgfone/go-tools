// Package handler is the handler collection of the logger.
package handler

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"sync"
	"time"

	"github.com/xgfone/go-tools/file"
)

const (
	// DAY_FMT is the date-rotaed format.
	DAY_FMT = "2006-01-02"

	// FILE_MODE is the mode to open the log file.
	FILE_MODE = os.O_APPEND | os.O_CREATE | os.O_WRONLY

	// FILE_PERM is the default permission to open the log file.
	FILE_PERM os.FileMode = 0666
)

var (
	dayRE       = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(\.\w+)?$`)
	day   int64 = 3600 * 24

	time2fmt = map[int64]string{
		day: DAY_FMT,
	}

	filePerm = FILE_PERM
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
func NewTimedRotatingFile(filename string) *TimedRotatingFile {
	filename, _ = filepath.Abs(filename)
	t := TimedRotatingFile{filename: filename, when: day, extRE: dayRE}
	t.SetBackupCount(31).SetInterval(1)
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

// SetBackupCount sets the number of the backup file. The default is 31.
func (t *TimedRotatingFile) SetBackupCount(num int) *TimedRotatingFile {
	t.backupCount = num
	return t
}

// SetInterval sets the interval day number to rotate. The default is 1.
func (t *TimedRotatingFile) SetInterval(interval int) *TimedRotatingFile {
	t.interval = int64(interval) * t.when
	t.reComputeRollover()
	return t
}

func (t *TimedRotatingFile) shouldRollover() bool {
	if time.Now().Unix() >= t.rotatorAt {
		return true
	}
	return false
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
	file, err := os.OpenFile(t.filename, FILE_MODE, FILE_PERM)
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
	fileNames, err := file.ListDir(dirName)
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
