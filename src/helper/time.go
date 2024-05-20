package helper

import "time"

func TimeISO8601() time.Time {
	timeNowStr := time.Now().Format(time.RFC3339)
	timeNow, err := time.Parse(time.RFC3339, timeNowStr)
	ErrorIfPanic(err)
	return timeNow
}
