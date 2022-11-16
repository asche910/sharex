package util

import (
	"fmt"
	"os/exec"
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
