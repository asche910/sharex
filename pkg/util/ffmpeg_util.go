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

	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("ffmpeg err ->", err)
	//	return
	//}
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ffmpeg output err ->", err)
		fmt.Println(string(bytes))

		return
	}
	fmt.Println(string(bytes))
}

// GetFirstFrame
// note: output file should be .jpg
func GetFirstFrame(input, output string) {
	RunFfmpegWithArgs("-ss", "1", "-i", input, "-vframes", "1", "-f", "image2", output)
}

//util.RunFfmpegWithArgs("-ss", "1", "-i", "/Users/as_/test.mp4", "-vframes", "1", "-f", "image2", "test.jpg")
