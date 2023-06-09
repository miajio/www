package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/bin"
	"github.com/miajio/www/lib"
)

type msgRouterImpl struct{}

type msgRouter bin.Router

var MsgRouter msgRouter = (*msgRouterImpl)(nil)

func (*msgRouterImpl) Running(e *gin.Engine) {
	msgGroup := e.Group("/msg")

	msgGroup.POST("/leave", func(ctx *gin.Context) {
		leave_mobile := ctx.DefaultPostForm("mobile", "")
		leave_name := ctx.DefaultPostForm("name", "")
		leave_msg := ctx.DefaultPostForm("msg", "")

		if leave_mobile == "" || leave_name == "" || leave_msg == "" {
			ctx.HTML(http.StatusOK, "error.html", gin.H{"code": 403, "msg": "参数错误", "err": "参数错误"})
			return
		}

		insertSQL := "INSERT INTO `leave_message` (`id`, `leave_mobile`, `leave_name`, `leave_msg`, `create_time`, `status`) VALUES (?, ?, ?, ?, NOW(), 1)"
		_, err := lib.DB.Exec(insertSQL, bin.UID(), leave_mobile, leave_name, leave_msg)
		if err != nil {
			ctx.HTML(http.StatusOK, "error.html", gin.H{"code": 500, "msg": "系统错误", "err": err.Error()})
			return
		}

		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
