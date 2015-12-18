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

func ToTime(t time.Time) string {
	return t.Format(FMT_TIME)
}

func ToDate(t time.Time) string {
	return t.Format(FMT_DATE)
}

func ToDateTime(t time.Time) string {
	return t.Format(FMT_DATETIME)
}
