package main

import (
	"fmt"
	"github.com/asche910/sharex/pkg/conf"
	"github.com/asche910/sharex/pkg/controller"
	"github.com/gin-gonic/gin"
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

	r.GET("/", controller.HomeController)
	r.GET("/download", controller.FileController)
	r.GET("/img/:name", controller.DownloadController)

	r.GET("/json", func(context *gin.Context) {
		context.JSON(200, data)
	})

	TEST()
	fmt.Println("listening: http://localhost:" + config["Port"])
	_ = r.Run(fmt.Sprintf(":%s", config["Port"]))
}

func TEST() {

	fmt.Println("---------------------------- TEST ----------------------------")

	fmt.Println("---------------------------- END! ----------------------------")
}
