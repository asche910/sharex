package controller

import (
	"fmt"
	"github.com/asche910/sharex/pkg/util"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
)

func HomeController(context *gin.Context) {
	loc, ok := context.GetQuery("loc")
	if !ok {
		loc = "."
	}

	newDir := filepath.Clean(loc)
	fmt.Println(loc, " --- ", newDir)
	if strings.Compare(loc, newDir) != 0 {
		context.Status(302)
		context.Header("Location", "/?loc="+newDir)
		return
	}

	files := util.GetShareXFiles(loc)
	context.HTML(200, "home.html", gin.H{
		"Files": files,
	})
}
