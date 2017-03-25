// Package datetime supplies some convenient functions about datetime.
package datetime

import (
	"time"
)

const (
	// DateFMT is the date format.
	DateFMT = "2006-01-02"

	// TimeFMT is the time format.
	TimeFMT = "15:04:05"

	// DateTimeFMT is the datetime format.
	DateTimeFMT = "2006-01-02 15:04:05"
)

// Now return the current unixstamp.
func Now() int64 {
	return time.Now().Unix()
}

// ToTime converts time.Time to "%H:%M:%S".
func ToTime(t time.Time) string {
	return t.Format(TimeFMT)
}

// ToDate converts time.Time to "%y-%m-%d".
func ToDate(t time.Time) string {
	return t.Format(DateFMT)
}

// ToDateTime converts time.Time to "%y-%m-%d %H:%M:%S".
func ToDateTime(t time.Time) string {
	return t.Format(DateTimeFMT)
}
