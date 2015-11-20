package utils

const (
	FMT_DATE     = "2006-01-02"
	FMT_DATETIME = "2006-01-02 15-04-05"
)

func Now() int64 {
	return time.Now().Unix()
}
