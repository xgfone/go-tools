// The handler of the logger.
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
	DAY_FMT = "2006-01-02"

	FILE_MODE = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	FILE_PERM = os.ModePerm
)

var (
	dayRE       = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(\.\w+)?$`)
	day   int64 = 3600 * 24

	time2fmt = map[int64]string{
		day: DAY_FMT,
	}
)

var (
	ErrFileNotOpen = errors.New("The file is not opened")
)

// A file handler based on the timed rotating, like
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

// Create a new TimedRotatingFile.
func NewTimedRotatingFile(filename string) *TimedRotatingFile {
	filename, _ = filepath.Abs(filename)
	t := TimedRotatingFile{filename: filename, when: day, extRE: dayRE}
	t.SetBackupCount(31).SetInterval(1)
	if err := t.open(); err != nil {
		panic(err)
	}
	return &t
}

// Write the string data into the file, which may rotate the file if necessary.
func (self *TimedRotatingFile) WriteString(data string) (n int, err error) {
	return self.Write([]byte(data))
}

// Write the byte slice data into the file, which may rotate the file if necessary.
func (self *TimedRotatingFile) Write(data []byte) (n int, err error) {
	self.Lock()
	defer self.Unlock()

	if self.w == nil {
		err = ErrFileNotOpen
		return
	}

	if self.shouldRollover() {
		if err = self.doRollover(); err != nil {
			return
		}
	}

	return self.w.Write(data)
}

// Set the number of the backup file. The default is 31.
func (self *TimedRotatingFile) SetBackupCount(num int) *TimedRotatingFile {
	self.backupCount = num
	return self
}

// Set the interval day number to rotate. The default is 1.
func (self *TimedRotatingFile) SetInterval(interval int) *TimedRotatingFile {
	self.interval = int64(interval) * self.when
	self.reComputeRollover()
	return self
}

func (self TimedRotatingFile) shouldRollover() bool {
	if time.Now().Unix() >= self.rotatorAt {
		return true
	}
	return false
}

func (self *TimedRotatingFile) Close() (err error) {
	if err = self.w.Close(); err != nil {
		return
	}
	self.w = nil
	return
}

func (self *TimedRotatingFile) open() error {
	file, err := os.OpenFile(self.filename, FILE_MODE, FILE_PERM)
	if err != nil {
		return err
	} else {
		self.w = file
		return nil
	}
}

func (self *TimedRotatingFile) doRollover() (err error) {
	if err = self.Close(); err != nil {
		return
	}

	dstTime := self.rotatorAt - self.interval
	dstPath := self.filename + "." + time.Unix(dstTime, 0).Format(time2fmt[self.when])
	if file.IsExist(dstPath) {
		os.Remove(dstPath)
	}

	if file.IsFile(self.filename) {
		if err = os.Rename(self.filename, dstPath); err != nil {
			return err
		}
	}

	if self.backupCount > 0 {
		for _, file := range self.getFilesToDelete() {
			os.Remove(file)
		}
	}

	self.reComputeRollover()
	return self.open()
}

func (self TimedRotatingFile) getFilesToDelete() []string {
	result := make([]string, 0, 30)
	dirName, baseName := filepath.Split(self.filename)
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
			if self.extRE.MatchString(suffix) {
				result = append(result, filepath.Join(dirName, fileName))
			}
		}
	}

	if len(result) <= self.backupCount {
		return []string{}
	} else {
		sort.Strings(result)
		return result[:len(result)-self.backupCount]
	}
}

func (self *TimedRotatingFile) reComputeRollover() {
	current_time := time.Now().Unix()

	t := time.Unix(current_time, 0)
	current_hour := t.Hour()
	current_minute := t.Minute()
	current_second := t.Second()

	r := self.interval - int64((current_hour*60+current_minute)*60+current_second)
	self.rotatorAt = current_time + r
}
