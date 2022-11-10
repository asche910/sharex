package util

import (
	"fmt"
	sharex "github.com/asche910/sharex/pkg"
	"os"
)

func init() {
	//fmt.Println("util init.")
}

func GetFiles(dir string) []string {
	var subFiles = []string{
		//".",
		"..",
	}

	dirs, _ := os.ReadDir(dir)

	for _, dir := range dirs {
		//fmt.Println(dir, dir.IsDir())
		subFiles = append(subFiles, dir.Name())
	}
	return subFiles
}

func GetShareXFiles(dir string) []sharex.ShareXFile {
	var retFiles = []sharex.ShareXFile{
		{IsFile: false, Name: "."},
		{IsFile: false, Name: ".."},
	}

	dirs, _ := os.ReadDir(dir)

	for _, dir := range dirs {

		fmt.Println()
		cur := sharex.ShareXFile{
			IsFile: !dir.IsDir(),
			Name:   dir.Name(),
		}
		if !dir.IsDir() {
			info, _ := dir.Info()
			//fmt.Println(info.Mode())
			cur.Size = info.Size()
		}
		retFiles = append(retFiles, cur)
	}
	return retFiles
}
