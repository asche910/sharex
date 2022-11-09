package main

import (
	"fmt"
	"github.com/asche910/sharex/pkg/conf"
	"github.com/asche910/sharex/pkg/util"
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

	files := util.GetFiles()

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "home.html", gin.H{
			"Dirs": files,
		})
	})

	r.GET("/json", func(context *gin.Context) {
		//context.HTML(200, "home.html")
		context.JSON(200, data)
	})

	TEST()
	fmt.Println("listening: http://localhost:" + config["Port"])
	_ = r.Run(fmt.Sprintf(":%s", config["Port"]))
}

func TEST() {

	fmt.Println("---------------------------- TEST ----------------------------")
	//stooges := []string{"Moe", "Larry", "Curly"} // len(stooges) == 3

	files := util.GetFiles()
	fmt.Println(files)
	fmt.Println("---------------------------- END! ----------------------------")
}
