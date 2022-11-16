package util

import (
	"fmt"
	"strings"
)

// GetFirstFrame
// note: output file should be .jpg
func GetFirstFrame(input, output string) {
	RunFfmpegWithArgs("-ss", "1", "-i", input, "-vframes", "1", "-f", "image2", output)
}

// GetMiddleFrame
// note: output file should be .jpg
// if failed, using GetFirstFrame
func GetMiddleFrame(input, output string) {
	videoLength, ok := GetVideoLength(input)
	if ok {
		midLen, ok := GetMiddleTime(videoLength)
		if ok {
			RunFfmpegWithArgs("-ss", midLen, "-i", input, "-vframes", "1", "-f", "image2", output)
			return
		}
	}
	RunFfmpegWithArgs("-ss", "1", "-i", input, "-vframes", "1", "-f", "image2", output)
}

// GetVideoLength
// example time: HH:mm:ss:-- 00:00:09:89
func GetVideoLength(input string) (string, bool) {
	resp, ok := RunFFprobeWithArgs(input)
	if !ok {
		return "", false
	}
	idx := strings.Index(resp, "Duration:")
	if idx < 0 {
		return "", false
	}
	timeLen := resp[idx+10 : idx+21]
	fmt.Println(timeLen)
	return timeLen, true
}
