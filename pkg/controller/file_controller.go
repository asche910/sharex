package controller

import (
	"fmt"
	"github.com/asche910/sharex/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func FileController(context *gin.Context) {
	fullName, ok := context.GetQuery("name")
	if !ok {
		//loc = "."
		context.JSON(http.StatusNotFound, nil)
		return
	}
	fullName = filepath.Clean(fullName)
	singleName := filepath.Base(fullName)

	f, err := os.Open(fullName)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	defer f.Close()
	fmt.Println("download:", fullName, singleName)

	if strings.LastIndex(singleName, ".") > 0 {
		context.File(fullName)
	} else {
		contentType, _ := util.GetFileContentType(f)
		fmt.Println("content-type", contentType)

		if strings.Compare("application/octet-stream", contentType) == 0 {
			fmt.Println("start download...", singleName)
			context.FileAttachment(fullName, singleName)
		} else {
			context.File(fullName)
		}
	}
}

func DownloadController(context *gin.Context) {
	fileName := context.Param("name")
	fullPath := filepath.Join(util.CacheDir, fileName)
	context.File(fullPath)
}
