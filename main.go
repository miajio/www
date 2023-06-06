package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/app"
	"github.com/miajio/www/bin"
)

var (
	// GinRouters gin路由集
	GinRouters bin.Routers = []bin.Router{
		app.UserRouter, // 用户路由
	}
)

func main() {
	e := gin.Default()

	e.Static("/static/images", "./static/images")
	e.Static("/static/js", "./static/js")

	bin.GinUtil.LoadHTMLFolders(e, []string{"./static"}, ".html")
	bin.GinUtil.LoadRouters(e, GinRouters...)

	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	})

	e.GET("/:page", func(ctx *gin.Context) {
		page := ctx.Param("page")
		if strings.Contains(page, ".html") {
			ctx.HTML(http.StatusOK, page, nil)
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "page is not found",
		})
	})

	e.Run(":8080")
}
