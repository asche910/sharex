package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sharex/src"
	"strings"
)
import "github.com/gin-gonic/gin"

// sharex
func main() {
	fmt.Println("sharex start")

	config := conf.InitConfig("sharex.conf")

	fmt.Println(config)
	r := gin.Default()

	var htmlFiles []string
	_ = filepath.Walk("web/view", func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
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

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "home.html", nil)
	})

	r.GET("/json", func(context *gin.Context) {
		//context.HTML(200, "home.html")
		context.JSON(200, data)
	})

	TEST()
	r.Run(fmt.Sprintf(":%s", config["Port"]))
}

func TEST() {

	fmt.Println("---------------------------- TEST ----------------------------")
	//stooges := []string{"Moe", "Larry", "Curly"} // len(stooges) == 3

	fmt.Println("---------------------------- END! ----------------------------")
}
