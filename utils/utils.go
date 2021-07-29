package utils

import (
	"strconv"
	"time"
)

func GetTimeFromUnix(dt int64) time.Time {

	i, err := strconv.ParseInt(strconv.Itoa(int(dt)), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}
