package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miajio/www/bin"
	"github.com/miajio/www/lib"
	"github.com/miajio/www/req"
)

type msgRouterImpl struct{}

type msgRouter bin.Router

var MsgRouter msgRouter = (*msgRouterImpl)(nil)

func (*msgRouterImpl) Running(e *gin.Engine) {
	msgGroup := e.Group("/msg")

	msgGroup.POST("/leave", func(ctx *gin.Context) {
		request := req.MsgLeaveRequest{}
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.HTML(http.StatusOK, "error.html", gin.H{"code": 403, "msg": "参数错误", "error": lib.TransError(err)})
			return
		}

		insertSQL := "INSERT INTO `leave_message` (`id`, `leave_mobile`, `leave_name`, `leave_msg`, `create_time`, `status`) VALUES (?, ?, ?, ?, NOW(), 1)"
		_, err := lib.DB.Exec(insertSQL, lib.UID(), request.Mobile, request.Name, request.Msg)
		if err != nil {
			ctx.HTML(http.StatusOK, "error.html", gin.H{"code": 500, "msg": "系统错误", "error": err.Error()})
			return
		}

		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
