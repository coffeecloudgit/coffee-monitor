package util

import (
	"time"
)

// StringToTime 2024-07-23T18:32:40.022+0800
func StringToTime(stringTime string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.999999999-0700" // Go的time包规定的日期格式字符串
	return time.Parse(layout, stringTime)
}
