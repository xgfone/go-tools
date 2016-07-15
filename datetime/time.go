// Some convenient functions about datetime.
package datetime

import (
	"time"
)

const (
	FMT_DATE     = "2006-01-02"
	FMT_DATETIME = "2006-01-02 15:04:05"
	FMT_TIME     = "15:04:05"
)

func Now() int64 {
	return time.Now().Unix()
}

// Convert time.Time to "%H:%M:%S".
func ToTime(t time.Time) string {
	return t.Format(FMT_TIME)
}

// Convert time.Time to "%y-%m-%d".
func ToDate(t time.Time) string {
	return t.Format(FMT_DATE)
}

// Convert time.Time to "%y-%m-%d %H:%M:%S".
func ToDateTime(t time.Time) string {
	return t.Format(FMT_DATETIME)
}
