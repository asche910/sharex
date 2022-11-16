package util

import (
	"fmt"
	"strconv"
)

// GetMiddleTime
// 02:04:08 -> 01:02:04.00
// 2:04:08:84 -> 01:02:04.42
func GetMiddleTime(timeLen string) (string, bool) {
	hour, err := strconv.Atoi(timeLen[0:2])
	if err != nil {
		return "", false
	}
	min, err := strconv.Atoi(timeLen[3:5])
	if err != nil {
		return "", false
	}
	sec, err := strconv.Atoi(timeLen[6:8])
	if err != nil {
		return "", false
	}
	var millSec int
	if len(timeLen) > 8 {
		atoi, err := strconv.Atoi(timeLen[9:11])
		fmt.Println("atoi ", atoi)
		if err == nil {
			millSec = atoi
		}
	}

	totalSec := hour*3600 + min*60 + sec
	hasOne := 0
	if totalSec%2 == 1 {
		hasOne = 1
	}
	totalSec /= 2

	hour = totalSec / 3600
	totalSec -= hour * 3600
	min = totalSec / 60
	totalSec -= min * 60

	totalSec = totalSec*2 + hasOne
	totalMillSec := totalSec*100 + millSec
	totalMillSec /= 2

	sec = totalMillSec / 100
	totalMillSec -= sec * 100
	millSec = totalMillSec

	return fmt.Sprintf("%02d:%02d:%02d.%02d", hour, min, sec, millSec), true
}
