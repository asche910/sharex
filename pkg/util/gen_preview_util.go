package util

import (
	"fmt"
	"strings"
)

func GetVideoPreView(name string) (string, bool) {
	dotIdx := strings.LastIndex(name, ".")

	uniqueName := Path2UniqueName(name[:dotIdx])
	uniqueName = "img_" + uniqueName + ".png"

	if CheckFileExists(CacheDir + uniqueName) {
		return uniqueName, true
	} else {
		dotIdx := strings.LastIndex(uniqueName, ".")
		nameNoTail := uniqueName[:dotIdx]

		_, err := GetSnapshot(name, CacheDir+nameNoTail, 1)
		if err != nil {
			fmt.Println("gen pre failed", name, err)
			return "", false
		}
		return uniqueName, true
	}
}
