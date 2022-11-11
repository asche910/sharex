package util

import (
	"fmt"
	sharex "github.com/asche910/sharex/pkg"
	"net/http"
	"os"
	"os/user"
	"strings"

	"golang.org/x/exp/slices"
)

var VIDEO_TYPE_LIST = []string{"mp4", "avi", "mov", "wmv", "flv"}

var UserHome string
var CacheDir string

func init() {
	//fmt.Println("util init.")

	u, err := user.Current()
	if err != nil {
		fmt.Println("get user home err", err)
	}
	UserHome = u.HomeDir
	fmt.Println("user home", UserHome)

	CacheDir = UserHome + "/.sharex/"
	os.Mkdir(CacheDir, 0777)
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

func GetShareXFiles(dirStr string) []sharex.ShareXFile {
	var retFiles = []sharex.ShareXFile{
		//{IsFile: false, Name: "."},
		//{IsFile: false, Name: "..", PreURL: "../static/image/test-large.png"},
	}

	dirs, _ := os.ReadDir(dirStr)

	for _, dir := range dirs {

		cur := sharex.ShareXFile{
			IsFile: !dir.IsDir(),
			Name:   dir.Name(),
		}

		// is File not dir
		if !dir.IsDir() {
			info, _ := dir.Info()

			//fmt.Println(info.Mode())
			cur.Size = info.Size()

			// todo load preview
			//fullPath := dirStr + "/" + dir.Name()

			fullPath := dirStr + "/" + dir.Name()
			fmt.Println(fullPath)
			if CheckVideoType(fullPath) {
				view, ok := GetVideoPreView(fullPath)
				if ok {
					cur.PreURL = "/img/" + view
				}
			}

		}
		fmt.Println(cur)
		retFiles = append(retFiles, cur)
	}
	return retFiles
}

func GetFileContentType(out *os.File) (string, error) {
	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func Path2UniqueName(path string) string {
	return strings.ReplaceAll(path, "/", "_")
}

func CheckVideoType(name string) bool {
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return false
	}
	fileType := strings.ToLower(name[idx+1:])

	isVideo := slices.Contains(VIDEO_TYPE_LIST, fileType)
	return isVideo
}

func CheckFileExists(name string) bool {
	_, err := os.Stat(name)
	fmt.Println(err)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
	//return !errors.Is(err, os.ErrNotExist)
}
