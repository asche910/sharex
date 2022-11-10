package main

import (
	"fmt"
	"github.com/asche910/sharex/pkg/conf"
	"github.com/asche910/sharex/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"

	"io/fs"
	"path/filepath"
	"strings"
)

// sharex
func main() {
	fmt.Println("sharex start")

	config := conf.InitConfig("sharex.conf")

	fmt.Println("conf:", config)
	r := gin.Default()

	_ = r.SetTrustedProxies(nil)
	r.Static("/static", "./web/static")

	var htmlFiles []string

	_ = filepath.WalkDir("web/view", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".html") {
			htmlFiles = append(htmlFiles, path)
		}
		return nil
	})
	fmt.Println("Load html files:", htmlFiles)
	r.LoadHTMLFiles(htmlFiles...)

	data := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{200, "success", true}

	//fmt.Println(data)

	r.GET("/", HomeController)
	r.GET("/download", DownloadController)

	r.GET("/json", func(context *gin.Context) {
		//context.HTML(200, "home.html")
		context.JSON(200, data)
	})

	TEST()
	fmt.Println("listening: http://localhost:" + config["Port"])
	_ = r.Run(fmt.Sprintf(":%s", config["Port"]))
}

func HomeController(context *gin.Context) {
	loc, ok := context.GetQuery("loc")
	if !ok {
		loc = "."
	}
	loc += "/"
	newDir := filepath.Dir(loc)
	fmt.Println(loc, " --- ", newDir)
	if strings.Contains(loc, "..") {
		context.Status(302)
		context.Header("Location", "/?loc="+newDir)
		return
	}

	files := util.GetShareXFiles(loc)
	context.HTML(200, "home.html", gin.H{
		"Files": files,
	})
}

func DownloadController(context *gin.Context) {
	fullName, ok := context.GetQuery("name")
	if !ok {
		//loc = "."
		context.JSON(http.StatusNotFound, nil)
	}

	idx := strings.LastIndex(fullName, "/")
	singleName := fullName[idx+1:]
	fmt.Println("download:", fullName, singleName)
	context.File(fullName)
	//context.FileAttachment(fullName, singleName)
}

func TEST() {

	fmt.Println("---------------------------- TEST ----------------------------")

	_, err := util.GetSnapshot("/Users/as_/Movies/giphy.mp4", "test", 1)
	if err != nil {
		return
	}

	fmt.Println("---------------------------- END! ----------------------------")
}
