package main

import (
	"context"
	"fmt"
	"github.com/asche910/sharex/pkg/conf"
	"github.com/asche910/sharex/pkg/controller"
	"github.com/gin-gonic/gin"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
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
		//context.HTML(200, "home.html")
		context.JSON(200, data)
	})

	TEST()
	fmt.Println("listening: http://localhost:" + config["Port"])
	_ = r.Run(fmt.Sprintf(":%s", config["Port"]))
}

func TEST() {

	fmt.Println("---------------------------- TEST ----------------------------")

	//fmt.Println(util.CheckFileExists("/Users/as_/Movies/giphy.mp4"))

	//util.GetPictureSnapshot("/Users/as_/go/sharex/web/static/image/test-large.png", "/Users/as_/go/sharex/web/static/image/test.png")

	ctx := context.Background()
	var cancel func()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	//exec.

	//cmd := exec.CommandContext(ctx, "ffprobe", "-l", "-a")
	////cmd := exec.Command("ffmpeg", "sharex.conf")
	//buf := bytes.NewBuffer(nil)
	//errBuf := bytes.NewBuffer(nil)
	//cmd.Stdout = buf
	//cmd.Stderr = errBuf
	//err := cmd.Run()
	////cmd.Ou
	//if err != nil {
	//	fmt.Println(err)
	//	//return "", err
	//}
	//output, _ := cmd.Output()
	//fmt.Println(string(output))

	//fmt.Println(string(buf.Bytes()))
	//fmt.Println(string(errBuf.Bytes()))

	//util.RunFfmpegWithArgs("-ss", "1", "-i", "/Users/as_/test.mp4", "-vframes", "1", "-f", "image2", "test.jpg")

	//util.GetFirstFrame("/Users/as_/test.mp4", "test.jpg")
	fmt.Println("---------------------------- END! ----------------------------")
}
