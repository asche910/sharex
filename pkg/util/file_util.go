package util

import (
	"fmt"
	sharex "github.com/asche910/sharex/pkg"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var VIDEO_TYPE_LIST = []string{"mp4", "avi", "mov", "wmv", "flv", "mkv"}
var PICTURE_TYPE_LIST = []string{"jpg", "jpeg", "png", "gif", "bmp", "svg", "webp"}

var UserHome string
var CacheDir string

func init() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("get user home err", err)
	}
	UserHome = u.HomeDir
	fmt.Println("user home", UserHome)

	CacheDir = fmt.Sprintf("%s%c.sharex", UserHome, os.PathSeparator)
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
		{IsFile: false, Name: "..", PreURL: "../static/image/test-large.png"},
	}

	isCacheDir := false
	absDir, _ := filepath.Abs(dirStr)
	if strings.Compare(absDir, CacheDir) == 0 {
		fmt.Println("visit cache dir!")
		isCacheDir = true
	}

	dirs, _ := os.ReadDir(dirStr)

	for _, dir := range dirs {

		cur := sharex.ShareXFile{
			IsFile: !dir.IsDir(),
			Name:   dir.Name(),
		}

		// if is File not dir
		if !dir.IsDir() {
			info, _ := dir.Info()
			cur.Size = info.Size()

			if isCacheDir {
				cur.PreURL = "/img/" + dir.Name()
			} else {
				fullPath := filepath.Join(dirStr, dir.Name())
				view, ok := GetPreview(fullPath)
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
	hashUint := Hash(path)
	return strconv.FormatUint(uint64(hashUint), 10)
}

func CheckIsVideo(name string) bool {
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return false
	}
	fileType := strings.ToLower(name[idx+1:])

	isVideo := slices.Contains(VIDEO_TYPE_LIST, fileType)
	return isVideo
}

func CheckIsPicture(name string) bool {
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return false
	}
	fileType := strings.ToLower(name[idx+1:])

	isPic := slices.Contains(PICTURE_TYPE_LIST, fileType)
	return isPic
}

func CheckFileExists(name string) bool {
	_, err := os.Stat(name)
	fmt.Println(err)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
