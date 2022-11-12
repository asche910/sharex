package util

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

// GetPreview
//	input -> fullPath
//  output -> filename of CACHE_DIR, if ok
func GetPreview(name string) (string, bool) {
	dotIdx := strings.LastIndex(name, ".")
	if dotIdx == -1 {
		return "", false
	}
	fileType := strings.ToLower(name[dotIdx+1:])
	if slices.Contains(VIDEO_TYPE_LIST, fileType) {
		return getVideoPreview(name)
	} else if slices.Contains(PICTURE_TYPE_LIST, fileType) {
		return getPicturePreview(name)
	}
	return "", false
}

func getVideoPreview(name string) (string, bool) {
	dotIdx := strings.LastIndex(name, ".")

	uniqueName := Path2UniqueName(name[:dotIdx])
	uniqueName = "img_" + uniqueName + ".png"

	if CheckFileExists(CacheDir + uniqueName) {
		return uniqueName, true
	} else {
		dotIdx := strings.LastIndex(uniqueName, ".")
		nameNoTail := uniqueName[:dotIdx]

		_, err := GetVideoSnapshot(name, CacheDir+nameNoTail, 1)
		if err != nil {
			fmt.Println("gen pre failed", name, err)
			return "", false
		}
		return uniqueName, true
	}
}

func getPicturePreview(name string) (string, bool) {
	dotIdx := strings.LastIndex(name, ".")

	uniqueName := Path2UniqueName(name[:dotIdx])
	uniqueName = "img_" + uniqueName + ".png"
	ok := GetPictureSnapshot(name, CacheDir+uniqueName)
	if !ok {
		return "", false
	}
	return uniqueName, true
}
