package models

import (
	"time"
)

/**
 * @Author wangdt
 * @Description //TODO ${end}
 * @Date 2022-09-15 19:34:58
 * @Param timestamp
 * @return string
 **/
func UnixTotime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0).UTC()
	return t.Format("2006-01-02 15:04:05")
}
