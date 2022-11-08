package util

import (
	"fmt"
	"os"
)

func __Init__() {

}

func GetFiles() []string {
	var subFiles []string

	dirs, _ := os.ReadDir(".")

	for _, dir := range dirs {
		fmt.Println(dir, dir.IsDir())
		subFiles = append(subFiles, dir.Name());
	}
	return subFiles
}
