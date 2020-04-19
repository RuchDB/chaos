package util

import (
	"fmt"
	"time"
)

func Now() time.Time {
	return time.Now()
}

// Timestamp in ms
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

// Full time format: 2020-01-01 00:00:00.0000
func FormatTimeFull(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d.%04d", 
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond() / 100000)
}

// Standard time format: 2020-01-01 00:00:00
func FormatTime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", 
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// Standard date format: 2020-01-01
func FormatDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
