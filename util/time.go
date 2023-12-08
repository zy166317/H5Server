package util

import "time"

// GetZeroClock 获取当天0点时间戳
func GetZeroClock() int64 {
	timeStr := time.Now().Format("2006-01-02")
	location, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return location.Unix()
}
