package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sharex/src"
	"strings"
)
import "github.com/gin-gonic/gin"

//
// sharex
//
func main() {
	fmt.Println("sharex start")
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
	fmt.Println(path)

	config := conf.InitConfig("justlike.conf")

	fmt.Println(config)
	r := gin.Default()

	data := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{200, "success", true}

	fmt.Println(data)
	r.GET("/", func(context *gin.Context) {
		//context.HTML(200, "home.html")
		context.JSON(200, data)
	})

	r.Run(fmt.Sprintf(":%s", config["Port"]))
}
