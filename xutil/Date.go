package xutil

import (
	"time"
)

func GetDate64(t ut_tv) time.Time {
	iTimeSec := int64(t.Tv_sec)
	sTimeSec := time.Unix(iTimeSec, 0)
	return sTimeSec
}