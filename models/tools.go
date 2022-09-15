package models

import (
	"time"
)

func UnixTotime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0).UTC()
	return t.Format("2006-01-02 15:04:05")
}
