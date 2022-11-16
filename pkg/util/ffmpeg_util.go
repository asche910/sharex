package util

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func RunFfmpegWithArgs(args ...string) {
	baseArgs := []string{"-hide_banner", "-loglevel", "error", "-y"}
	args = append(baseArgs, args...)
	fmt.Println(args)
	cmd := exec.Command("ffmpeg", args...)

	bytes, err := cmd.CombinedOutput()
	resp := string(bytes)
	if err != nil {
		fmt.Println("ffmpeg output err ->", err, resp)
		return
	}
	fmt.Println(resp)
}

func RunFFprobeWithArgs(args ...string) (string, bool) {
	baseArgs := []string{"-hide_banner"}
	args = append(baseArgs, args...)
	fmt.Println(args)
	cmd := exec.Command("ffprobe", args...)

	bytes, err := cmd.CombinedOutput()
	resp := string(bytes)
	if err != nil {
		fmt.Println("ffmpeg exec err ->", err, resp)
		return "", false
	}
	fmt.Println(resp)
	return resp, true
}

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
