package util

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"os"
)

func GetPictureSnapshot(path, outPath string) bool {
	imgData, _ := os.ReadFile(path)
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//生成缩略图，尺寸150*200，并保持到为文件2.jpg
	image = imaging.Resize(image, 100, 0, imaging.Lanczos)
	err = imaging.Save(image, outPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
