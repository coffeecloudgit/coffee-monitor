package util

import (
	"strings"
	"time"
)

// StringToTime 2024-07-23T18:32:40.022+0800
func StringToTime(stringTime string) (time.Time, error) {
	if strings.HasSuffix(stringTime, "Z") {
		return time.Parse(time.RFC3339, stringTime)
	}
	return time.Parse("2006-01-02T15:04:05.999999999-0700", stringTime)
}
func GetTimeOfDay(expireBlockHeight, currentBlockHeight uint64) string {
	var blockNum int64 = 0
	blockNum = int64(expireBlockHeight - currentBlockHeight)
	expireUnix := time.Now().Unix() + blockNum*30
	expireTime := time.Unix(expireUnix, 0)
	return expireTime.Format("2006-01-02")
}
