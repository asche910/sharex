package util

import (
	"os"
)

func init() {
	//fmt.Println("util init.")
}

func GetFiles() []string {
	var subFiles []string

	dirs, _ := os.ReadDir(".")

	for _, dir := range dirs {
		//fmt.Println(dir, dir.IsDir())
		subFiles = append(subFiles, dir.Name())
	}
	return subFiles
}
